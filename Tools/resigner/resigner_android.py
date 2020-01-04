#!/usr/bin/python
#coding=utf-8

import os
import json
import shutil
import zipfile

config = None

pathPackage = ""
pathInput = ""

def json_load_byteified(file_handle):
  return _byteify(
    json.load(file_handle, object_hook=_byteify),
    ignore_dicts=True
  )

def json_loads_byteified(json_text):
  return _byteify(
    json.loads(json_text, object_hook=_byteify),
    ignore_dicts=True
  )

def _byteify(data, ignore_dicts = False):
  if isinstance(data, unicode):
    return data.encode('utf-8')
  if isinstance(data, list):
    return [ _byteify(item, ignore_dicts=True) for item in data ]
  if isinstance(data, dict) and not ignore_dicts:
    return {
      _byteify(key, ignore_dicts=True): _byteify(value, ignore_dicts=True)
      for key, value in data.iteritems()
    }
  return data

def loadConfig():
  global config
  f = open('config.json')
  config = json_load_byteified(f)

def walk2work():
  global pathInput
  global pathOut

  pathInput = os.path.join(os.getcwd(), "input")
  pathOut = os.path.join(os.getcwd(), "output")

  for f in os.listdir(pathInput):
    pathChild = os.path.join(pathInput, f)
    pathextractall = os.path.join(pathOut, f)
    pkgname = "%s_%s_t.apk"%(config["package_name"], f)
    pkgpath = os.path.join(pathOut, pkgname)

    if os.path.isdir(pathChild):
      extractall(pathextractall) #解压安装包
      replacefiles(pathChild, pathextractall) #文件替换
      compressall(pathextractall, pkgpath) #生成压缩包
      resignpackage(pkgpath, f) #重签安装包
      shutil.rmtree(pathextractall) #删除解压文件夹

def replacefiles(pathsrc, pathdst):
  pathcfg = os.path.join(pathsrc, 'apk_config.json')
  if os.path.exists(pathcfg):
    f = open(pathcfg)
    configresign = json_load_byteified(f)
    for fname in configresign["files_remove"]:
      fpath = os.path.join(pathdst, fname)
      if os.path.exists(fpath):
        os.remove(fpath)

    for pname in configresign["directorys_remove"]:
      dpath = os.path.join(pathdst, pname)
      if os.path.exists(dpath):
        shutil.rmtree(dpath)

  pathJoins = os.path.join(pathsrc, "apk")
  for root, dirs, files in os.walk(pathJoins, topdown=True):
    for dirname in dirs:
      dirsrc = os.path.join(root, dirname)
      dirdst = dirsrc.replace(pathJoins, pathdst)

      if not os.path.exists(dirdst):
        os.makedirs(dirdst)

    for fname in files:
      fsrc = os.path.join(root, fname)
      fdst = fsrc.replace(pathJoins, pathdst)
      shutil.copy(fsrc, fdst)

def zip_dir(dirname, zipfilename):
  filelist = []
  if os.path.isfile(dirname):
    filelist.append(dirname)
  else :
    for root, dirs, files in os.walk(dirname):
      for name in files:
        filelist.append(os.path.join(root, name))
  zf = zipfile.ZipFile(zipfilename, "w", zipfile.zlib.DEFLATED)
  for tar in filelist:
    arcname = tar[len(dirname):]
    zf.write(tar,arcname)
  zf.close()

def zip_append(dirname, zipfilename, pathDelta):
  filelist = []
  if os.path.isfile(dirname):
    filelist.append(dirname)
  else :
    for root, dirs, files in os.walk(dirname):
      for name in files:
        filelist.append(os.path.join(root, name))
  zf = zipfile.ZipFile(zipfilename, "a", zipfile.zlib.DEFLATED)
  for tar in filelist:
    arcname = tar[len(dirname)+1:]
    if pathDelta != "":
      arcname = os.path.join(pathDelta, arcname)
    zf.write(tar,arcname)
  zf.close()

def extractall(pathextract):
  global pathPackage

  pathPackage = os.path.join(os.getcwd(), "package", "%s.apk"%(config["package_name"]))
  zf = zipfile.ZipFile(pathPackage)
  zf.extractall(pathextract)
  zf.close()
  pathUnuse = os.path.join(pathextract, "META-INF")
  shutil.rmtree(pathUnuse)

def compressall(pathextractall, pkgpath):
  zip_dir(pathextractall, pkgpath)

def resignpackage(fpath, fn):
  pathKeystore = os.path.join(os.getcwd(), "package", config["keystore"])
  pathResigned = os.path.join(pathOut, "%s_%s.apk"%(config["package_name"], fn))

  cmd = "jarsigner -verbose -keystore %s -storepass %s -signedjar %s -digestalg SHA1 -sigalg MD5withRSA %s %s"%(pathKeystore, config["alias_password"], pathResigned, fpath, config["alias"])
  ret = os.system(cmd)

  if ret == 0:
    print u"版本[%s]重签成功!\n"%(pathResigned)
  else:
    print u"版本[%s]重签失败!\n"%(pathResigned)

  os.remove(fpath)

if __name__ == '__main__':
  loadConfig()
  walk2work()