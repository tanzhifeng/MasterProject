#!/usr/bin/python
#coding=utf-8

import os
import json
import PIL.Image

config = None

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
	f = open('assetsmaker.json')
	config = json_load_byteified(f)

def makeAssets():
	global config

	makes = config['makes']

	for value in makes:
		pathOrg = os.path.join(os.getcwd(), value['original'])
		pathTrg = os.path.join(os.getcwd(), value['output'])
		width = value['width']
		height = value['height']

		if not os.path.exists(pathTrg):
			os.makedirs(pathTrg)

		oimg = PIL.Image.open(pathOrg)
		filename = os.path.join(pathTrg, value['name'])
		nimg = oimg.resize((width, height), PIL.Image.ANTIALIAS)
		nimg.save(filename, "PNG")

if __name__ == '__main__':
	loadConfig()
	makeAssets()