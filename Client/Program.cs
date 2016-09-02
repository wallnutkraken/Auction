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




        void Run()
        {
            Auction auction = new Auction();

            while(true)
            {
                auction.DecideOnReplies();
                auction.DisplayMenu();
                int choice = int.Parse(Console.ReadLine());
                switch(choice)
                {
                    case 0:
                        auction.RequestUpdate();
                        break;
                    case 1:
                        auction.DisplayItemsToBidOn();
                        int[] choices = auction.AskWhichItemToBidOn();
                        auction.Bid(choices[0], choices[1]);
                        break;
                }
            }


        }
    }
}

