val i = 5 : Int
var j = 6
var mySeq = Seq[Any](1,3,4,5,6,7,4.5,2.3d,8l)
var array = Array(1,2,3,4,4)

def printList[T](list: Seq[T]): String =
    list match  {
    case head +: tail => s"$head " + printList(tail)
    case Nil => "Nil"
  }

println(printList(mySeq))
println(printList(array))

case class Result(votes:Int,voteShare:Double,partyCode:String)

def process(results:Seq[Result]): Seq[String] =
  results.map( (res) => s"{votes = ${res.votes}," +
    s"share: ${res.voteShare}, party = ${res.partyCode} }")

process(Seq(Result(1000,35.5,"LAB"),
  Result(800,30.1,"CON"),
  Result(400,15.5,"LD"))) foreach println
