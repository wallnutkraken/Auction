﻿using System;
using System.Collections.Concurrent;
using System.Net;
using System.Net.Mime;
using System.Net.Sockets;
using System.Text;
using Newtonsoft.Json;


namespace Client
{
    public class Client
    {
        public string IpRemote { get; set; }
        public int PortNr { get; set; }
        public BlockingCollection<SendCommand> CommandsToSend = new BlockingCollection<SendCommand>();
        public BlockingCollection<Reply> Replies = new BlockingCollection<Reply>();

        private Socket _socket;
        private byte[] Encode<T>(T objToEncode)
        {
            string objLiteral = JsonConvert.SerializeObject(objToEncode);
            byte[] objBytes = Encoding.UTF8.GetBytes(objLiteral);
            return objBytes;
        }
        private TResult Decode<TResult>(byte[] receivedBytes)
        {
            string receivedString = Encoding.UTF8.GetString(receivedBytes);
            TResult result = JsonConvert.DeserializeObject<TResult>(receivedString);
            return result;
        }

        public Client(string ipRemote, int portNr)
        {
            IpRemote = ipRemote;
            PortNr = portNr;
            _socket = new Socket(SocketType.Stream, ProtocolType.Tcp);
        }

        private void SendToRemote<T>(T objToSend)
        {
            byte[] message = Encode(objToSend);

            if (!_socket.Connected)
            {
                _socket.Connect(new IPEndPoint(IPAddress.Parse(IpRemote), PortNr));
            }

            _socket.Send(message);

        }

        public void SendCommands()
        {
            while (true)
            {
                SendCommand command;
                CommandsToSend.TryTake(out command);
                if (command != null)
                {
                    SendToRemote(command);
                }
            }
        }

        public void ListenToReplies()
        {
            while (true)
            {
                if (!_socket.Connected)
                {
                    _socket.Connect(new IPEndPoint(IPAddress.Parse(IpRemote), PortNr));
                }
                byte[] receiveBuffer = new byte[1024];
                int relevantLenght = _socket.Receive(receiveBuffer);
                byte[] relevantBuffer = new byte[1024];
                receiveBuffer.CopyTo(relevantBuffer, 0);
                Reply reply = Decode<Reply>(relevantBuffer);
                Replies.TryAdd(reply);

            }
        }

    }
}