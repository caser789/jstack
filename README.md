# jstack

1. 初始化 stack 的时候创建 endpoint? 还是 NewEndpoint 时创建 endpoint?
    - User create TransportEndpoint
    - User create LinkEndpoint
    - NetworkEndpoint 在 stack.AddAddress时创建
