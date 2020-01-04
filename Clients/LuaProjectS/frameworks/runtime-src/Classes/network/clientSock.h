#ifndef __HEADER_CLIENT_NET_
#define __HEADER_CLIENT_NET_

#include <queue>
#include <mutex>
#include <functional>
#ifdef WIN32
#include <WinSock2.h>
#include <ws2tcpip.h>
#include <mstcpip.h>
#else
#include <arpa/inet.h>
#include <netinet/tcp.h>
#endif

#define HEAD_LEN 2
#define PROTOC_LEN 4

enum ClientStatus {
	Succeed,
	Failed,
	Losted,
	Dataed
};

struct Packet {
	ClientStatus status;
	unsigned int ptid;
	char* content;
	unsigned short len;
};

class clientSock{
    
private:
    int m_sock;
    char m_ip[32];
    int m_port;
	
	bool connected;

	std::mutex mutexStatus;
	std::mutex mutexSend;
	std::mutex mutexRecv;

	std::queue<Packet*> queueSend;
	std::queue<Packet*> queueRecv;

	std::function<void(Packet* ret)> m_onReply;
public:
    clientSock(){
#ifdef WIN32
		WSADATA wsaData;
        WSAStartup(MAKEWORD(2, 2),&wsaData);
#endif
		connected = false;
    }

	void connectByIp(const char* ip, int port, std::function<void(Packet* ret)> reply);
	void connectByDomain(const char* domain, int port, std::function<void(Packet* ret)> reply);
    void closeSocket();

	void checkUpdate();

	void handlerSend(unsigned short ptid, char* data = NULL, unsigned short dataLen = 0);
protected:
	void push2QueueSend(Packet* p);
	void cleanQueueSend();

	void push2QueueRecv(Packet* p);
	void cleanQueueRecv();

	void start_thread();
	void send_thread();
	void recv_thread();

    int recv_len(char* buffer, int len);
	int send_len(char* buffer, int len);

	static char* getIpByDomain(const char* domain);
};


#endif
