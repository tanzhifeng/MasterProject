#!/usr/bin/python
#coding=utf-8

import os
import shutil
import sys, getopt
import platform

if __name__ == '__main__':
	try:
		opts, args = getopt.getopt(sys.argv[1:], "hi:o:", ["help","input=","output="])
	except getopt.GetoptError:
		print "usage: %s -i <input path> -o <output path>"%(sys.argv[0])
		sys.exit()

	for opt, arg in opts:
		if opt in ("-h", "--help"):
			print "usage: %s -i <input path> -o <output path>"%(sys.argv[0])
			print "       %s --input <input path> -output <output path>"%(sys.argv[0])
			sys.exit()
		elif opt in ("-i", "--input"):
			pathInput = arg
		elif opt in ("-o", "--output"):
			pathOutput = arg

	if not pathInput:
		print "usage: no input path"
		sys.exit()

	if not pathOutput:
		print "usage: no output path"
		sys.exit()

	commond = ""
	if platform.system() == 'Windows':
		commond = "protoc.exe --go_out={0} {1}"
	else:
		commond = "./protoc --go_out={0} {1}"

	for root, dirs, files in os.walk(pathInput, topdown=True):
		for dirname in dirs:
			dirout = os.path.join(pathOutput, root, dirname)
			if not os.path.exists(dirout):
				os.makedirs(dirout)

		for fname in files:
			ifile = os.path.join(root, fname)
			os.system(commond.format(pathOutput, ifile))