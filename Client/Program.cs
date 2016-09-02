using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using Newtonsoft.Json;

namespace Client
{
    class Program
    {
        static void Main(string[] args)
        {
            Program p = new Program();
            p.Run();
        }

        private TResult Decode<TResult>(string receivedJson)
        {
            TResult result = JsonConvert.DeserializeObject<TResult>(receivedJson);
            return result;
        }

        void Run()
        {
            Client client = new Client("127.0.0.1", 1632);
            SendCommand availableItemsToBidOn = new SendCommand("list", null);
            SendCommand bidCommand = new SendCommand("bid", new[] { "0", "120" });
            client.Run();

            client.CommandsToSend.Add(availableItemsToBidOn);
            client.CommandsToSend.Add(bidCommand);

            List<AuctionPimp> auctionPimps = new List<AuctionPimp>();

            while(true)
            {
                Reply reply = client.Replies.Take();
                Console.WriteLine(reply);

                string replyType = reply.ReplyType;
                switch(replyType)
                {
                    case "list":
                        List<AuctionPimp> aps = Decode<List<AuctionPimp>>(reply.ValueJson);
                        auctionPimps.AddRange(aps);
                        break;
                    case "somethingElse":
                        break;
                    default:
                        break;
                }
                foreach (AuctionPimp auctionPimp in auctionPimps)
                {
                    Console.WriteLine(auctionPimp);
                }
            }
        }
    }
}
