using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Sockets;
using System.Text;
using Newtonsoft.Json;

namespace ServerSharp
{
    public class Server
    {
        private Socket _socket;
        private IPEndPoint _localEndPoint;
        private int _portNr;
        private List<AuctionPimp> runningAuctions = new List<AuctionPimp>() { new AuctionPimp() };

        public Server(int portNr)
        {
            IPHostEntry ipaddresses = Dns.GetHostEntry("localhost");
            IPAddress thing = ipaddresses.AddressList[1];
            _localEndPoint = new IPEndPoint(thing, portNr);
            _socket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
        }


        public void Listen()
        {
            while(true)
            {
                if(!_socket.IsBound)
                {
                    _socket.Bind(_localEndPoint);
                    Console.WriteLine("Started listening ...");
                    _socket.Listen(10);
                    _socket = _socket.Accept();
                    Console.WriteLine("Client connected ...");
                    byte[] buffer = new byte[0xffff];
                    int bytesCount = _socket.Receive(buffer);

                    SendCommand receivedCommand = Decode<SendCommand>(buffer);

                    switch(receivedCommand.Command)
                    {
                        case "list":
                            RespondToList();
                            Console.WriteLine("Responding to list request");
                            break;
                        default:
                            break;
                    }

                    _socket.Disconnect(true);
                }
            }
        }

        private void Send<T>(T objectToSend)
        {
            byte[] message = Encode(objectToSend);
            _socket.Send(message);
        }

        private void RespondToList()
        {
            Send(new Reply("list", EncodeToString(runningAuctions)));
        }

        private byte[] Encode<T>(T objToEncode)
        {
            string objLiteral = JsonConvert.SerializeObject(objToEncode);
            byte[] objBytes = Encoding.UTF8.GetBytes(objLiteral);
            return objBytes;
        }

        private string EncodeToString<T>(T objectToEncode)
        {
            string objLiteral = JsonConvert.SerializeObject(objectToEncode);
            return objLiteral;
        }
        private TResult Decode<TResult>(byte[] receivedBytes)
        {
            string receivedString = Encoding.UTF8.GetString(receivedBytes);
            TResult result = JsonConvert.DeserializeObject<TResult>(receivedString);
            return result;
        }
    }


}