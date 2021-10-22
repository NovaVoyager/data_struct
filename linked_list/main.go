package main

import "fmt"

func main() {
	test()
}

func test() {
	linked := Constructor() //0
	linked.AddAtHead(38)    //1
	linked.AddAtTail(66)    //2
	linked.AddAtTail(61)    //3
	linked.AddAtTail(76)    //4
	linked.AddAtTail(26)    //5
	linked.AddAtTail(37)    //6
	linked.AddAtTail(8)     //7
	linked.DeleteAtIndex(5) //8
	linked.AddAtHead(4)     //9

	linked.AddAtHead(45)       //10
	fmt.Println(linked.Get(4)) //11
	linked.AddAtTail(85)       //12
	linked.AddAtHead(37)       //13
	fmt.Println(linked.Get(5)) //14
	linked.AddAtTail(93)       //15
	linked.AddAtIndex(10, 23)  //16
	linked.AddAtTail(21)       //17
	linked.AddAtHead(52)       //18
	linked.AddAtHead(15)       //19

	linked.AddAtHead(47)        //20
	fmt.Println(linked.Get(12)) //21
	linked.AddAtIndex(6, 24)    //22
	linked.AddAtHead(64)        //23
	fmt.Println(linked.Get(4))  //24
	linked.AddAtHead(31)        //25
	linked.DeleteAtIndex(6)     //26
	linked.AddAtHead(40)        //27
	linked.AddAtTail(17)        //28
	linked.AddAtTail(15)        //29

	linked.AddAtIndex(19, 2)    //30
	linked.AddAtTail(11)        //31
	linked.AddAtHead(86)        //32
	fmt.Println(linked.Get(17)) //33
	linked.AddAtTail(55)        //34
	linked.DeleteAtIndex(15)    //35
	linked.AddAtIndex(14, 95)   //36
	linked.DeleteAtIndex(22)    //37
	linked.AddAtHead(66)        //38
	linked.AddAtTail(95)        //39

	linked.AddAtHead(8)         //40
	linked.AddAtHead(47)        //41
	linked.AddAtTail(23)        //42
	linked.AddAtTail(39)        //43
	fmt.Println(linked.Get(30)) //44
	fmt.Println(linked.Get(27)) //45
	linked.AddAtHead(0)         //46
	linked.AddAtTail(99)        //47
	linked.AddAtTail(45)        //48
	linked.AddAtTail(4)         //49

	linked.AddAtIndex(9, 11)   //50
	fmt.Println(linked.Get(6)) //51
	linked.AddAtHead(81)       //52
	linked.AddAtIndex(18, 32)  //53
	linked.AddAtHead(20)       //54
	linked.AddAtTail(13)       //55
	linked.AddAtTail(42)       //56
	linked.AddAtIndex(37, 91)  //57
	linked.DeleteAtIndex(36)   //58
	linked.AddAtIndex(10, 37)  //59

	linked.AddAtHead(96)       //60
	linked.AddAtHead(57)       //61
	linked.DeleteAtIndex(20)   //62
	linked.AddAtTail(89)       //63
	linked.DeleteAtIndex(18)   //64
	linked.AddAtIndex(41, 5)   //65
	linked.AddAtTail(23)       //66
	linked.AddAtHead(75)       //67
	fmt.Println(linked.Get(7)) //68
	linked.AddAtIndex(25, 51)  //69

	linked.AddAtTail(48)        //70
	linked.AddAtHead(46)        //71
	linked.AddAtHead(29)        //72
	linked.AddAtHead(85)        //73
	linked.AddAtHead(82)        //74
	linked.AddAtHead(6)         //75
	linked.AddAtHead(38)        //76
	linked.DeleteAtIndex(14)    //77
	fmt.Println(linked.Get(1))  //78
	fmt.Println(linked.Get(12)) //79

	linked.AddAtHead(42)        //80
	fmt.Println(linked.Get(42)) //81
	linked.AddAtTail(83)        //82
	linked.AddAtTail(13)        //83
	linked.AddAtIndex(14, 20)   //84
	linked.AddAtIndex(17, 34)   //85
	linked.AddAtHead(36)        //86
	linked.AddAtTail(58)        //87
	linked.AddAtTail(2)         //88
	fmt.Println(linked.Get(38)) //89

	linked.AddAtIndex(33, 59)   //90
	linked.AddAtHead(37)        //91
	linked.DeleteAtIndex(15)    //92
	linked.AddAtTail(64)        //93
	fmt.Println(linked.Get(56)) //94
	linked.AddAtHead(0)         //95
	fmt.Println(linked.Get(40)) //96
	linked.AddAtHead(92)        //97
	linked.DeleteAtIndex(63)    //98
	fmt.Println(linked.Get(35)) //99

	linked.AddAtTail(62) //100
	linked.AddAtTail(32) //101

	fmt.Println(linked)
}
