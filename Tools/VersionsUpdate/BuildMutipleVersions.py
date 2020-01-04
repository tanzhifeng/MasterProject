# -*- coding: utf-8-*-
import json
import pysvn
import os
import shutil
import xxhash
import zipfile
from hashlib import md5

config = None
client = None

def getFileMd5(filepath):
  m = md5()
  f = open(filepath, 'rb')
  m.update(f.read())
  f.close()
  return m.hexdigest()

def getHashFilename(filename):
	global config
	xxhashEnable = config["xxhashEnable"]
	if xxhashEnable:
		xxhashSeed = config["xxhashSeed"]
		return str(xxhash.xxh32(filename, xxhashSeed).intdigest())
	else:
		return filename

def get_login( realm, username, may_save ):
	global config
	return True, config["username"], config["password"], False

def checkUpdate():
	global config
	global client

	client = pysvn.Client()
	client.callback_get_login = get_login

	sources = config["sources"]
	support64bit = config["support64bit"]
	export = config["export"]

	if not os.path.exists(sources):
		os.makedirs(sources)

	if os.path.exists(export):
		shutil.rmtree(export)
	os.makedirs(export)

	modules = config["modules"]
	for name in modules:
		module = modules[name]
		moduleVersions = module["versions"]
		moduleSvn = module["svnUrl"]
		versionTop = moduleVersions[-1]
		rvTop = pysvn.Revision(pysvn.opt_revision_kind.number, versionTop)
		pathLocal = "%s/%s"%(sources, name)
		if os.path.exists(pathLocal):
			print "[%s] Start update module"%(name)
			client.update(pathLocal, revision = rvTop)
		else:
			print "[%s] Start checkout module"%(name)
			client.checkout(moduleSvn, pathLocal, revision = rvTop)

		moduleCompile(name)
		moduleDifference(name)
		moduleJson(name)

def moduleCompile(modulename):
	print "[%s] Start compile module"%(modulename)

	secretKey = config["secretKey"]
	secretSign = config["secretSign"]
	sources = config["sources"]
	support64bit = config["support64bit"]
	pathLocal = "%s/%s"%(sources, modulename)

	pathSrc = "%s/src"%(pathLocal)
	pathDst = "%s/out/src"%(pathLocal)
	commond = "cocos3172 luacompile -s %s -d %s -e -k %s -b %s>nul"%(pathSrc, pathDst, secretKey, secretSign)
	os.system(commond)

	if support64bit:
		commond = "cocos3172 luacompile -s %s -d %s -e -k %s -b %s --bytecode-64bit>nul"%(pathSrc, pathDst, secretKey, secretSign)
		os.system(commond)

	pathRes = "%s/res"%(pathLocal)
	pathDst = "%s/out/res"%(pathLocal)
	if os.path.exists(pathDst):
		shutil.rmtree(pathDst)
	shutil.copytree(pathRes, pathDst)

def moduleDifference(modulename):
	print "[%s] Start difference module"%(modulename)

	global config
	global client

	export = config["export"]
	sources = config["sources"]
	modules = config["modules"]
	exthold = config["exthold"]
	support64bit = config["support64bit"]

	module = modules[modulename]
	moduleVersions = module["versions"]
	moduleSvn = module["svnUrl"]
	versionTop = moduleVersions[-1]
	rvTop = pysvn.Revision(pysvn.opt_revision_kind.number, versionTop)
	pathLocal = "%s/%s"%(sources, modulename)

	for version in moduleVersions:
		pathSrc = "%s/%s/out"%(sources, modulename)
		pathDst = "%s/%s/%d/%s"%(export, modulename, version, modulename)

		if version == versionTop:
			shutil.copytree(pathSrc, pathDst)
		else:
			summary = client.diff_summarize(moduleSvn, pysvn.Revision(pysvn.opt_revision_kind.number, version), moduleSvn, rvTop)

			for item in summary:
				if item.node_kind == pysvn.node_kind.file and item.summarize_kind != pysvn.diff_summarize_kind.delete:
					itemPath = item.path
					(filePath, filename) = os.path.split(item.path)
					(shotname, extension) = os.path.splitext(filename)
					if extension == ".lua":
						itemPath = item.path.replace('.lua','.luac')
					pathFileSrc = "%s/%s"%(pathSrc, itemPath)
					pathFileDst = "%s/%s"%(pathDst, itemPath)

					fullPath = "%s/%s"%(pathDst, filePath)
					if not os.path.exists(fullPath):
						os.makedirs(fullPath)
					shutil.copyfile(pathFileSrc, pathFileDst)

					if support64bit and extension == ".lua":
						itemPath = item.path.replace('.lua','_64.luac')
						pathFileSrc = "%s/%s"%(pathSrc, itemPath)
						pathFileDst = "%s/%s"%(pathDst, itemPath)

						fullPath = "%s/%s"%(pathDst, filePath)
						if not os.path.exists(fullPath):
							os.makedirs(fullPath)
						shutil.copyfile(pathFileSrc, pathFileDst)

		pathLoop = "%s/%s/%d"%(export, modulename, version)

		for rt, dirs, files in os.walk(pathLoop, topdown=False):
			for f in files:
				(shotname, extension) = os.path.splitext(f)
				nf = getHashFilename(f)
				if extension in exthold:
					nf = "%s%s"%(nf, extension)
				os.rename(os.path.join(rt, f), os.path.join(rt, nf))

			for d in dirs:
				nd = getHashFilename(d)
				os.rename(os.path.join(rt, d), os.path.join(rt, nd))

		pathZip = "%s/%s/%d-%d.zip"%(export, modulename, version, versionTop)
		directoryCompress(pathLoop, pathZip)
		shutil.rmtree(pathLoop)

def directoryCompress(dirname, zipfilename):
	filelist = []
	if os.path.isfile(dirname):
		filelist.append(dirname)
	else:
		for root, dirs, files in os.walk(dirname):
			for name in files:
				filelist.append(os.path.join(root, name))

	zf = zipfile.ZipFile(zipfilename, "w", zipfile.zlib.DEFLATED)
	for tar in filelist:
		arcname = tar[len(dirname):]
		zf.write(tar,arcname)
	zf.close()

def moduleJson(modulename):
	print "[%s] Start module json"%(modulename)

	global config

	export = config["export"]
	modules = config["modules"]
	module = modules[modulename]
	moduleVersions = module["versions"]
	versionTop = moduleVersions[-1]

	jsonDic = {}
	jsonDic["updates"] = {}
	jsonDic["version"] = versionTop
	updates = jsonDic["updates"]

	for version in moduleVersions:
		package = "%d-%d.zip"%(version, versionTop)
		filename = "%s/%s/%s"%(export, modulename, package)

		update = {}
		update["version"] = versionTop
		update["size"] = os.path.getsize(filename)
		update["package"] = package
		update["md5"] = getFileMd5(filename)
		updates[version] = update

	nameJson = "%s/%s/%s.json"%(export, modulename, modulename)
	fd = open(nameJson, 'w')
	json.dump(jsonDic, fd)
	fd.close()

def byteify(input):
	if isinstance(input, dict):
		return {byteify(key):byteify(value) for key,value in input.iteritems()}
	elif isinstance(input, list):
		return [byteify(element) for element in input]
	elif isinstance(input, unicode):
		return input.encode('utf-8')
	else:
		return input

def loadConfig():
	global config

	f = open("config.json", 'r')
	config = byteify(json.load(f))

if __name__=="__main__":
	loadConfig()
	checkUpdate()