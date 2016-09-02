using System;
using System.Collections.Generic;
using Newtonsoft.Json;

namespace Client
{
    public class Auction
    {
        public List<AuctionPimp> RunningAuctionPimps;
        private Client _client;

        public Auction()
        {
            _client = new Client("127.0.0.1", 1632);
            _client.Run();
            RunningAuctionPimps = new List<AuctionPimp>();
        }

        public void Bid(int auctionPimpId, int biddingPrice)
        {
            SendCommand bidCommand = new SendCommand("bid", new[] { auctionPimpId.ToString(), biddingPrice.ToString() });
            _client.CommandsToSend.Add(bidCommand);
        }

        public void RequestUpdate()
        {
            SendCommand requestRunningAuctionPimps = new SendCommand("list", null);
            _client.CommandsToSend.Add(requestRunningAuctionPimps);
        }

        private TResult Decode<TResult>(string receivedJson)
        {
            TResult result = JsonConvert.DeserializeObject<TResult>(receivedJson);
            return result;
        }

        public int[] AskWhichItemToBidOn()
        {
            Console.WriteLine("choose item");
            int itemchoice = int.Parse(Console.ReadLine());
            Console.WriteLine("gimme price");
            int price = int.Parse(Console.ReadLine());

            return new[] {itemchoice, price};
        }
        public void DecideOnReplies()
        {
            while(_client.Replies.Count > 0)
            {
                Reply reply = _client.Replies.Take();

                string replyType = reply.ReplyType;
                switch(replyType)
                {
                    case "list":
                        List<AuctionPimp> aps = Decode<List<AuctionPimp>>(reply.ValueJson);
                        RunningAuctionPimps = new List<AuctionPimp>(aps);
                        break;
                    default:
                        break;
                }
            }
        }

        public void DisplayMenu()
        {
            Console.WriteLine("menu");
            Console.WriteLine("0 ... update");
            Console.WriteLine("1 ... bid");
        }

        public void DisplayItemsToBidOn()
        {
            if(RunningAuctionPimps.Count > 0)
            {
                Console.WriteLine("items to bid on");

                for(int i = 0; i < RunningAuctionPimps.Count; i++)
                {
                    Console.WriteLine(
                        $"{i} .. AuctionPimpID:{RunningAuctionPimps[i].Id},ItemName:{RunningAuctionPimps[i].BidItem},CurrentPrice:{RunningAuctionPimps[i].CurrentBid}");
                }
            }
            else
            {
                Console.WriteLine("No items to bid on.");
            }
        }
    }












}