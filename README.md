# 即时通讯系统

#### 目标功能

注册、登陆、在线用户列表、群聊（广播）、点对点聊天、离线留言

#### 实现传输协议内容

消息协议：struct message={string type, string data='特定结构体类型实例实例化'}

发送方式：先长度，再发消息体本身

客户端接收长度，并根据长度接受内容并输出。
