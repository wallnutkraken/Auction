using System;
using System.Collections.Concurrent;
using System.Net;
using System.Net.Mime;
using System.Net.Sockets;
using System.Text;
using System.Threading;
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
        private bool _connected;
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
            try
            {
                _socket.Connect(new IPEndPoint(IPAddress.Parse(IpRemote), PortNr));
                _connected = true;
            }
            catch(Exception e)
            {
                Console.WriteLine(e);
            }
        }

        public void Run()
        {
            Thread sendingThread = new Thread(SendCommands);
            Thread receiveThread = new Thread(ListenToReplies);
            sendingThread.Start();
            receiveThread.Start();
        }

        public void Stop()
        {
            _connected = false;
        }

        private void SendToRemote<T>(T objToSend)
        {
            byte[] message = Encode(objToSend);
            _socket.Send(message);
        }

        public void SendCommands()
        {
            while(_connected)
            {
                SendCommand command;
                CommandsToSend.TryTake(out command);
                if(command != null)
                {
                    SendToRemote(command);
                }
            }
        }

        public void ListenToReplies()
        {
            while(_connected)
            {
                byte[] receiveBuffer = new byte[0xffff];
                int relevantLenght = _socket.Receive(receiveBuffer);
                byte[] relevantBuffer = new byte[relevantLenght];
                for(int i = 0; i < relevantLenght; i++)
                {
                    relevantBuffer[i] = receiveBuffer[i];
                }
                Reply reply = Decode<Reply>(relevantBuffer);
                Replies.TryAdd(reply);
            }
        }
    }
}