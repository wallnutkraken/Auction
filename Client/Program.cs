using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
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
            Client client = new Client("127.0.0.1",1632);
            SendCommand availableItemsToBidOn = new SendCommand("list", null);
            Reply allItems = client.SendToRemote<Reply, SendCommand>(availableItemsToBidOn);

            Console.WriteLine(allItems);

            SendCommand bidCommand = new SendCommand("bid",new []{"10","20"});
            Reply confirmation = client.SendToRemote<Reply, SendCommand>(bidCommand);

            Console.WriteLine(confirmation);

            Console.ReadLine();
        }
    }
}
