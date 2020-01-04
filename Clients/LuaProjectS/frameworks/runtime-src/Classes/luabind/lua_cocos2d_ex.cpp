#include "scripting/lua-bindings/manual/CCLuaEngine.h"
#include "scripting/lua-bindings/manual/tolua_fix.h"
#include "scripting/lua-bindings/manual/LuaBasicConversions.h"


#include "clientSock.h"
#include "cocos2d.h"
//#include "AppDelegate.h"
#if (CC_TARGET_PLATFORM==CC_PLATFORM_ANDROID)
#include "platform/android/jni/JniHelper.h"
#endif

static int tolua_util_makeTag(lua_State *L)
{
    int a = tolua_tonumber(L, 2, 0);
    int b = tolua_tonumber(L, 3, 0);
    
    int c = ((int)(((unsigned short)(((unsigned int)(a)) & 0xffff)) | ((unsigned int)((unsigned short)(((unsigned int)(b)) & 0xffff))) << 16));
    
    tolua_pushnumber(L, c);
    
    return 1;
}

static int tolua_util_makeVersion(lua_State* L)
{
	const int productVer = (const int)tolua_tonumber(L, 2, 0);
	const int mainVer = (const int)tolua_tonumber(L, 3, 0);
	const int subVer = (const int)tolua_tonumber(L, 4, 0);
	const int buildVer = (const int)tolua_tonumber(L, 5, 0);

	unsigned int ret = ((unsigned int)(((unsigned char)(productVer)) << 24) | (((unsigned char)(mainVer)) << 16) | (((unsigned char)(subVer)) << 8) | ((unsigned char)(buildVer)));
	tolua_pushnumber(L, ret);

	return 1;
}

#ifdef WIN32

#define LEN_NETWORK_ID 13

//状态信息
struct tagAstatInfo
{
	ADAPTER_STATUS				AdapterStatus;						//网卡状态
	NAME_BUFFER					NameBuff[16];						//名字缓冲
};

//网卡地址
bool getWinMachineId(char* szMachineId)
{
	//变量定义
	HINSTANCE hInstance = NULL;

	//执行逻辑
	__try
	{
		//加载 DLL
		hInstance = LoadLibrary(TEXT("netapi32.dll"));
		if (hInstance == NULL) __leave;

		//获取函数
		typedef BYTE __stdcall NetBiosProc(NCB * Ncb);
		NetBiosProc * pNetBiosProc = (NetBiosProc *)GetProcAddress(hInstance, "Netbios");
		if (pNetBiosProc == NULL) __leave;

		//变量定义
		NCB Ncb;
		LANA_ENUM LanaEnum;
		ZeroMemory(&Ncb, sizeof(Ncb));
		ZeroMemory(&LanaEnum, sizeof(LanaEnum));

		//枚举网卡
		Ncb.ncb_command = NCBENUM;
		Ncb.ncb_length = sizeof(LanaEnum);
		Ncb.ncb_buffer = (BYTE *)&LanaEnum;
		if ((pNetBiosProc(&Ncb) != NRC_GOODRET) || (LanaEnum.length == 0)) __leave;

		//获取地址
		if (LanaEnum.length>0)
		{
			//变量定义
			tagAstatInfo Adapter;
			ZeroMemory(&Adapter, sizeof(Adapter));

			//重置网卡
			Ncb.ncb_command = NCBRESET;
			Ncb.ncb_lana_num = LanaEnum.lana[0];
			if (pNetBiosProc(&Ncb) != NRC_GOODRET) __leave;

			//获取状态
			Ncb.ncb_command = NCBASTAT;
			Ncb.ncb_length = sizeof(Adapter);
			Ncb.ncb_buffer = (BYTE *)&Adapter;
			Ncb.ncb_lana_num = LanaEnum.lana[0];
			strcpy((char *)Ncb.ncb_callname, "*");
			if (pNetBiosProc(&Ncb) != NRC_GOODRET) __leave;

			//获取地址
			for (int i = 0; i < 6; i++)
			{
				assert((i * 2)<LEN_NETWORK_ID);
				sprintf(&szMachineId[i * 2], "%02X", Adapter.AdapterStatus.adapter_address[i]);
			}
		}
	}

	//结束清理
	__finally
	{
		//释放资源
		if (hInstance != NULL)
		{
			FreeLibrary(hInstance);
			hInstance = NULL;
		}

		//错误断言
		if (AbnormalTermination() == TRUE)
		{
			assert(FALSE);
		}
	}

	return true;
}

static int tolua_util_getMachineId(lua_State* L)
{
	char szMachineId[255] = { 0 };
	bool ret = getWinMachineId(szMachineId);
	
	tolua_pushstring(L, szMachineId);
	return 1;
}

#endif

static int tolua_clientsock_new(lua_State *L)
{
	clientSock* self = new clientSock();
	tolua_pushusertype(L, (void*)self, "clientSock");

	return 1;
}

static int tolua_clientsock_connectByIp(lua_State *L)
{
	clientSock* self = (clientSock*)tolua_tousertype(L, 1, NULL);
	const char* ip = (const char*)tolua_tostring(L, 2, "");
	const short port = (const short)tolua_tonumber(L, 3, 0);
	cocos2d::LUA_FUNCTION handlerReply = (toluafix_ref_function(L, 4, 0));

	self->connectByIp(ip, port, [=](Packet* packet) {
		tolua_pushnumber(L, packet->status);
		tolua_pushnumber(L, packet->ptid);
		lua_pushlstring(L, packet->content, packet->len);
		LuaEngine::getInstance()->getLuaStack()->executeFunctionByHandler(handlerReply, 3);
	});
	return 0;
}

static int tolua_clientsock_connectByDomain(lua_State *L)
{
	clientSock* self = (clientSock*)tolua_tousertype(L, 1, NULL);
	const char* domain = (const char*)tolua_tostring(L, 2, "");
	const short port = (const short)tolua_tonumber(L, 3, 0);
	cocos2d::LUA_FUNCTION handlerReply = (toluafix_ref_function(L, 4, 0));

	self->connectByDomain(domain, port, [=](Packet* packet) {
		tolua_pushnumber(L, packet->status);
		tolua_pushnumber(L, packet->ptid);
		lua_pushlstring(L, packet->content, packet->len);
		LuaEngine::getInstance()->getLuaStack()->executeFunctionByHandler(handlerReply, 3);
	});
	return 0;
}

static int tolua_clientsock_closeSocket(lua_State *L)
{
	clientSock* self = (clientSock*)tolua_tousertype(L, 1, NULL);
	self->closeSocket();

	return 0;
}

static int tolua_clientsock_checkUpdate(lua_State *L)
{
	clientSock* self = (clientSock*)tolua_tousertype(L, 1, NULL);
	self->checkUpdate();

	return 0;
}

static int tolua_clientsock_handlerSend(lua_State *L)
{
	clientSock* self = (clientSock*)tolua_tousertype(L, 1, NULL);
	const unsigned short ptid = (const unsigned short)tolua_tonumber(L, 2, 0);

	int argc = lua_gettop(L) - 1;

	if (argc < 2)
		self->handlerSend(ptid);
	else
	{
		size_t len = 0;
		char* content = (char*)lua_tolstring(L, 3, &len);
		char* data = (char*)malloc(len);
		memcpy(data, content, len);
		self->handlerSend(ptid, data, len);
	}

	return 0;
}

static int lua_register_cocosd_ex(lua_State* tolua_S)
{
	std::string typeName;

//    tolua_open(tolua_S);
//    tolua_module(tolua_S, "util", 0);
//    tolua_beginmodule(tolua_S,"util");
//    tolua_function( tolua_S, "makeTag", tolua_util_makeTag );
//    tolua_function(tolua_S, "createQRSprite", tolua_util_createQRSprite);
//    tolua_function(tolua_S, "saveQRImageFile", tolua_util_saveQRImageFile);
//    tolua_function(tolua_S, "makeVersion", tolua_util_makeVersion);
//#ifdef WIN32
//    tolua_function(tolua_S, "getMachineId", tolua_util_getMachineId);
//#endif
//    tolua_endmodule(tolua_S);

	tolua_usertype(tolua_S, "clientSock");
	tolua_cclass(tolua_S, "clientSock", "clientSock", "clientSock", NULL);
	tolua_beginmodule(tolua_S, "clientSock");
	tolua_function(tolua_S, "new", tolua_clientsock_new);
	tolua_function(tolua_S, "connectByIp", tolua_clientsock_connectByIp);
	tolua_function(tolua_S, "connectByDomain", tolua_clientsock_connectByDomain);
	tolua_function(tolua_S, "closeSocket", tolua_clientsock_closeSocket);
	tolua_function(tolua_S, "checkUpdate", tolua_clientsock_checkUpdate);
	tolua_function(tolua_S, "handlerSend", tolua_clientsock_handlerSend);
	tolua_constant(tolua_S, "Succeed", ClientStatus::Succeed);
	tolua_constant(tolua_S, "Failed", ClientStatus::Failed);
	tolua_constant(tolua_S, "Losted", ClientStatus::Losted);
	tolua_constant(tolua_S, "Dataed", ClientStatus::Dataed);
	tolua_endmodule(tolua_S);
	typeName = typeid(clientSock).name();
	g_luaType[typeName] = "clientSock";
	g_typeCast["clientSock"] = "clientSock";


    return 1;
}

extern int register_cocosd_ex(lua_State* L)
{
    lua_getglobal(L, "_G");
    if (lua_istable(L,-1))//stack:...,_G,
    {
        lua_register_cocosd_ex(L);
    }
    lua_pop(L, 1);
    return 1;
}
