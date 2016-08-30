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
            Client client = new Client("1.1.1.1",12345);

            string response = client.SendToRemote<string,string>("hello");
        }
    }
}
