namespace ServerSharp
{
    public class Reply
    {
        public string ReplyType { get; set; }
        public string ValueJson { get; set; }

        public Reply()
        {
            
        }

        public Reply(string replyType, string valueJson)
        {
            ReplyType = replyType;
            ValueJson = valueJson;
        }
        public override string ToString()
        {
            return $"{ReplyType},{ValueJson}";
        }
    }
}