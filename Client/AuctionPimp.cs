using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Client
{
    public class AuctionPimp
    {
        public int Id { get; set; }
        public int StartPrice { get; set; }
        public int ExpDate { get; set; }
        public Item BidItem { get; set; }
        public int CurrentBid { get; set; }
    }
}
