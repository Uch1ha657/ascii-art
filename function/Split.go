package function


func Split(ascii string)[][]string {
	slise:= []string{}
	lines := [][]string{}
	line := ""
	for _, v := range ascii {
		if v == '\n'{
			 slise = append (slise,line)
			line = ""

		}else{
			line+= string(v)
		}
		
	}
	if line != "" {

		 slise = append (slise,line)
		line = ""
		
	}
	for i := 1; i < len(slise); i+=8 {
		if i+8 >len(slise){
			break
		}
		abed:= slise[i:i+8]
		lines= append(lines, abed)
		i++
	}
	return lines

}