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
        public TResult SendToRemote<TResult, TO>(TO objToSend)
        {
            byte[] message = Encode(objToSend);
            _socket.Connect(new IPEndPoint(IPAddress.Parse(IpRemote), PortNr));
            _socket.Send(message);
            byte[] receiveBuffer = new byte[1024];
            int relevantLenght = _socket.Receive(receiveBuffer);
            byte[] relevantBuffer = new byte[relevantLenght];
            receiveBuffer.CopyTo(relevantBuffer, 0);
            TResult result = Decode<TResult>(relevantBuffer);
            return result;
        }
    }
}