package main

import "fmt"

func ganho(t [9]string)int{//0=O ganhou, 1=X ganhou, 2=velha, 3=ainda há casas vazias
	//verticais O
	for i:=0; i<3; i++{
		if string(t[i])=="O" && string(t[i+3])=="O" && string(t[i+6])=="O"{
			return 0;
		}
	}
	//verticais X
	for i:=0; i<3; i++{
		if string(t[i])=="X" && string(t[i+3])=="X" && string(t[i+6])=="X"{
			return 1;
		}
	}
	//horizontais O
	for i:=0; i<9; i+=3{
		if string(t[i])=="O" && string(t[i+1])=="O" && string(t[i+2])=="O"{
			return 0;
		}
	}
	//horizontais X
	for i:=0; i<9; i+=3{
		if string(t[i])=="X" && string(t[i+1])=="X" && string(t[i+2])=="X"{
			return 1;
		}
	}
	//diagonais O e X
	if string(t[0])=="O" && string(t[4])=="O" && string(t[8])=="O"{
		return 0;
	}else if string(t[2])=="O" && string(t[4])=="O" && string(t[6])=="O"{
		return 0;
	}else if string(t[0])=="X" && string(t[4])=="X" && string(t[8])=="X"{
		return 1;
	}else if string(t[2])=="X" && string(t[4])=="X" && string(t[6])=="X"{
		return 1;
	}
	//caso ninguém tenha ganho, confere se há casas vazias
	for i:=0; i<9; i++{
		if string(t[i])==" "{
			return 3;
		}
	}
	return 2;
}
func robo(t[9]string, c bool)int{//c==true=círculo, c==false=cruz
	var l string;
	var a string;
	if c{
		l="O";
		a="X";
	}else{
		l="X";
		a="O";
	}
	//verticais do robô
	for i:=0; i<3; i++{
		if string(t[i])==l&&string(t[i+3])==l&&string(t[i+6])==" "{
			return i+6;
		}
	}
	for i:=3; i<6;i++{
		if string(t[i])==l&&string(t[i+3])==l&&string(t[i-3])==" "{
			return i-3;
		}
	}
	for i:=6; i<9; i++{
		if string(t[i])==l&&string(t[i-6])==l&&string(t[i-3])==" "{
			return i-3;
		}
	}
	//horizontais do robô
	for i:=0; i<9; i+=3{
		if string(t[i])==l&&string(t[i+1])==l&&string(t[i+2])==" "{
			return i+2;
		}
	}
	for i:=1; i<9; i+=3{
		if string(t[i])==l&&string(t[i+1])==l&&string(t[i-1])==" "{
			return i-1;
		}
	}
	for i:=2; i<9; i+=3{
		if string(t[i])==l&&string(t[i-2])==l&&string(t[i-1])==" "{
			return i-1;
		}
	}
	//diagonais do robô(improgramáveis)
	if string(t[0])==l&&string(t[4])==l&&string(t[8])==" "{
		return 8;
	}else if string(t[4])==l&&string(t[8])==l&&string(t[0])==" "{
		return 0;
	}else if string(t[2])==l&&string(t[4])==l&&string(t[6])==" "{
		return 6;
	}else if string(t[6])==l&&string(t[4])==l&&string(t[2])==" "{
		return 2;
	}else if string(t[8])==l&&string(t[0])==l&&string(t[4])==" "{
		return 4;
	}else if string(t[6])==l&&string(t[2])==l&&string(t[4])==" "{
		return 4;
	}
	//se não houver uma vitória imediata para o robô, ele vai procurar uma vitória para o adversário
	//verticais do adversário
	for i:=0; i<3; i++{
		if string(t[i])==a&&string(t[i+3])==a&&string(t[i+6])==" "{
			return i+6;
		}
	}
	for i:=3; i<6;i++{
		if string(t[i])==a&&string(t[i+3])==a&&string(t[i-3])==" "{
			return i-3;
		}
	}
	for i:=6; i<9; i++{
		if string(t[i])==a&&string(t[i-6])==a&&string(t[i-3])==" "{
			return i-3;
		}
	}
	//horizontais do adversário
	for i:=0; i<9; i+=3{
		if string(t[i])==a&&string(t[i+1])==a&&string(t[i+2])==" "{
			return i+2;
		}
	}
	for i:=1; i<9; i+=3{
		if string(t[i])==a&&string(t[i+1])==a&&string(t[i-1])==" "{
			return i-1;
		}
	}
	for i:=2; i<9; i+=3{
		if string(t[i])==a&&string(t[i-2])==a&&string(t[i-1])==" "{
			return i-1;
		}
	}
	//diagonais do adversário(improgramáveis)
	if string(t[0])==a&&string(t[4])==a&&string(t[8])==" "{
		return 8;
	}else if string(t[4])==a&&string(t[8])==a&&string(t[0])==" "{
		return 0;
	}else if string(t[2])==a&&string(t[4])==a&&string(t[6])==" "{
		return 6;
	}else if string(t[6])==a&&string(t[4])==a&&string(t[2])==" "{
		return 2;
	}else if string(t[8])==a&&string(t[0])==a&&string(t[4])==" "{
		return 4;
	}else if string(t[6])==a&&string(t[2])==a&&string(t[4])==" "{
		return 4;
	}
	//se ainda não houver ameaças, ele devve criá-las, mas, como substituto, ele vai tentar ocupar as melhores casas
	if c{
		for i:=0; i<4; i+=2{
			if string(t[i])==" "{
				return i;
			}
		}
		if string(t[4])==" "{
			return 4;
		}
		for i:=6; i<9; i+=2{
			if string(t[i])==" "{
				return i;
			}
		}
		for i:=1; i<9; i+=2{
			if string(t[i])==" "{
				return i;
			}
		}
	}else{
		if string(t[4])==" "{
			return 4;
		}
		for i:=0; i<4; i+=2{
			if string(t[i])==" "{
				return i;
			}
		}
		for i:=6; i<9; i+=2{
			if string(t[i])==" "{
				return i;
			}
		}
		for i:=1; i<9; i+=2{
			if string(t[i])==" "{
				return i;
			}
		}
	}
	return 4;//se tudo der errado
}
func printtab(t[9]string){
	fmt.Println("---------");
	for i:=0; i<9;i++{
		if (i+1)%3==0{
			fmt.Printf("%s. \n", string(t[i]));
		}else{
			fmt.Printf("%s. ", string(t[i]));
		}
	}
	fmt.Println("---------");
}
func limpatab()[9]string{
	var t=[9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "};
	return t;
}

func main(){
	var tab[9]string=limpatab();
	var l int;
	var quem int;//quem começa: eu=1; robô=2
	pergunta:
		fmt.Println("Quer começar? Se sim, digite 1; se não, digite 2.");
		fmt.Scanln(&quem);
	if quem!=1&&quem!=2{
		goto pergunta;
	}
	if quem==1{
		for i:=0; ganho(tab)==3; i++{
			printtab(tab);
			if i%2==0{
				jogo:
					fmt.Scanln(&l);//casa do topo da esquerda é 0
				if tab[l]!=" " ||l>=9{
					fmt.Println("Essa casa não está disponível.");
					goto jogo;
				} else{
					tab[l]="O";
				}
			}else{
				tab[robo(tab, false)]="X";
			}
		}
		if ganho(tab)==2{
			fmt.Println("Deu velha.");
		} else if ganho(tab)==0{
			fmt.Println("Círculo ganhou.");
		} else{
			fmt.Println("Cruz ganhou.");
		}
	}else{
		for i:=0; ganho(tab)==3; i++{
			printtab(tab);
			if i%2!=0{
				perg:
					fmt.Scanln(&l);//casa do topo da esquerda é 0
				if tab[l]!=" " ||l>=9{
					fmt.Println("Essa casa não está disponível.");
					goto perg;
				} else{
					tab[l]="X";
				}
			}else{
				tab[robo(tab, true)]="O";
			}
		}
		if ganho(tab)==2{
			fmt.Println("Deu velha.");
		} else if ganho(tab)==0{
			fmt.Println("Círculo ganhou.");
		} else{
			fmt.Println("Cruz ganhou.");
		}
		printtab(tab);
	}
	/*
	jogo entre duas pessoas
	for i:=0; ganho(tab)==3; i++{
		printtab(tab);
		jogo:
			fmt.Scanln(&l);//casa do topo da esquerda é 0
			if tab[l]!=" " ||l>=9{
				fmt.Println("Essa casa não está disponível.");
				goto jogo;
			} else{
				if i%2==0{
					tab[l]="O";
				} else if i%2!=0{
					tab[l]="X";
				}
			}
	}
	if ganho(tab)==2{
		fmt.Println("Deu velha.");
	} else if ganho(tab)==0{
		fmt.Println("Círculo ganhou.");
	} else{
		fmt.Println("Cruz ganhou.");
	}
	*/
}