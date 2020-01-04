import os,sys
cwd = os.getcwd()
ResPath = os.path.dirname(cwd)
APPPath = ResPath+"/runtime/mac/LuaProjectS-desktop.app/Contents/MacOS/LuaProjectS-desktop"
if os.path.exists(APPPath):
	os.system(APPPath)
else:
	print("not exists !")
