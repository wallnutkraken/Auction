namespace ServerSharp
{
    public class AuctionPimp
    {

        public int Id { get; set; }
        public int StartPrice { get; set; }
        public long ExpDate { get; set; }
        public Item BidItem { get; set; }
        public int CurrentBid { get; set; }
        public override string ToString()
        {
            return $"{Id},{StartPrice},{ExpDate},{BidItem},{CurrentBid}";
        }
    }

    
}