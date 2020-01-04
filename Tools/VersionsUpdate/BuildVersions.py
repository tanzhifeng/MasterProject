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

	print u"开始更新模块..."

	client = pysvn.Client()
	client.callback_get_login = get_login

	sources = config["sources"]
	support64bit = config["support64bit"]

	if not os.path.exists(sources):
		os.makedirs(sources)

	modules = config["modules"]
	for name in modules:
		versions = modules[name]
		lastVersion = versions[-1]
		rv = pysvn.Revision(pysvn.opt_revision_kind.number, lastVersion)
		localPath = "%s/%s"%(sources, name)
		svnUrl = "%s/%s"%(config["svnUrl"], name)

		if os.path.exists(localPath):
			client.update(localPath, revision=rv)
		else:
			client.checkout(svnUrl, localPath, revision=rv)

		pathSrc = "%s/src"%(localPath)
		pathDst = "%s/out/src"%(localPath)
		commond = "cocos3172 luacompile -s %s -d %s -e -k %s -b %s>nul"%(pathSrc, pathDst, config["secretKey"], config["secretSign"])
		os.system(commond)

		if support64bit:
			commond = "cocos3172 luacompile -s %s -d %s -e -k %s -b %s --bytecode-64bit>nul"%(pathSrc, pathDst, config["secretKey"], config["secretSign"])
			os.system(commond)

		pathRes = "%s/res"%(localPath)
		pathDst = "%s/out/res"%(localPath)
		if os.path.exists(pathDst):
			shutil.rmtree(pathDst)
		shutil.copytree(pathRes, pathDst)

def zip_dir(dirname,zipfilename):
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

def checkDifference():
	global config
	global client

	print u"开始版本差异比较..."

	export = config["export"]
	if os.path.exists(export):
		shutil.rmtree(export)
	os.makedirs(export)

	sources = config["sources"]
	modules = config["modules"]
	exthold = config["exthold"]
	xxhashEnable = config["xxhashEnable"]
	xxhashSeed = config["xxhashSeed"]
	support64bit = config["support64bit"]
	for name in modules:
		versions = modules[name]
		lastVersion = versions[-1]
		rv = pysvn.Revision(pysvn.opt_revision_kind.number, lastVersion)
		localPath = "%s/%s"%(sources, name)
		svnUrl = "%s/%s"%(config["svnUrl"], name)

		for indexV in range(len(versions)):
			version = versions[indexV]
			pathSrc = "%s/%s/out"%(sources, name)
			pathDst = "%s/%s/%d/%s"%(export, name, version, name)

			if version == lastVersion:
				shutil.copytree(pathSrc, pathDst)
			else:
				summary = client.diff_summarize(svnUrl, pysvn.Revision(pysvn.opt_revision_kind.number, version), svnUrl, rv)
				for indexS in range(len(summary)):
					item = summary[indexS]
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


			pathLoop = "%s/%s/%d"%(export, name, version)
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

			pathZip = "%s/%s/%d-%d.zip"%(export, name, version, lastVersion)
			if version == lastVersion:
				pathZip = "%s/%s/%d.zip"%(export, name, lastVersion)
			zip_dir(pathLoop, pathZip)
			shutil.rmtree(pathLoop)

def checkJson():
	global config

	print u"开始导出JSON文件..."

	export = config["export"]
	modules = config["modules"]

	for name in modules:
		jsonDic = {}
		jsonDic["updates"] = {}
		
		versions = modules[name]
		lastVersion = versions[-1]
		jsonDic["version"] = lastVersion
		updates = jsonDic["updates"]

		for indexV in range(len(versions)):
			version = versions[indexV]
			package = "%d-%d.zip"%(version, lastVersion)
			if version == lastVersion:
				package = "%d.zip"%(lastVersion)
			filename = "%s/%s/%s"%(export, name, package)

			update = {}
			update["version"] = lastVersion
			update["size"] = os.path.getsize(filename)
			update["package"] = package
			update["md5"] = getFileMd5(filename)
			updates[version] = update

		nameJson = "%s/%s/%s.json"%(export, name, name)
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

	print u"读取配置..."
	f = open("config.json", 'r')
	config = byteify(json.load(f))

if __name__=="__main__":
	loadConfig()
	checkUpdate()
	checkDifference()
	checkJson()