using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace Client
{
    class Program
    {
        static void Main(string[] args)
        {
            Program p = new Program();
            p.Run();
        }

        void Run()
        {
            Client client = new Client("127.0.0.1", 1632);
            SendCommand availableItemsToBidOn = new SendCommand("list", null);
            SendCommand bidCommand = new SendCommand("bid", new[] { "10", "20" });


            Thread sendingThread = new Thread(client.SendCommands);
            Thread receiveThread = new Thread(client.ListenToReplies);
            sendingThread.Start();
            receiveThread.Start();

            client.CommandsToSend.Add(availableItemsToBidOn);
            client.CommandsToSend.Add(bidCommand);

            while (true)
            {
                Reply reply = client.Replies.Take();
                Console.WriteLine(reply);
            }

        }
    }
}
