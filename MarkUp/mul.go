package mul

type prop struct{Name string; Value string}
type Tag struct{
	Name,InnerText string
	Props []prop
	Children []Tag
	Type byte 
}

type MULDS struct{Data []Tag}

func getProps(src string)[]prop{
	n, v, q, Props := "", "", 0, make([]prop,0)
	for _,c:=range src {
		if c=='"'  {q+=1}
		if q==1 && c!='"' {v+=string(c)}
		if q==2 {q, n = 0, ""}
		if q==0 && c!='=' && c!='/'  && c>32 {n+=string(c)}
		if q==0 && len(n)>0 && (c=='=' || c<=32) {
			if len(v)>0 {Props[len(Props)-1].Value, v, n = v, "", ""; continue}
			Props, n = append(Props,prop{Name:n}), ""
		}
	}
 return Props
}

func strip(src string)string{
	e := 0
	for i:=len(src)-1; src[i]!='<' && i>0; i-- {e=i}
	return src[:len(src)-(e+1)]
}

func printProps(Props []prop)string{
	res, prop := "", ""
	for _,p:=range Props {
		prop = " "+p.Name+` = "`+p.Value+`"`
		if len(p.Value)==0 {prop = " "+p.Name}
		res+= prop
	}
	return res
}

func Parse(src *string) []Tag {
	var currentTag, lastTag, content, Props string
	n, p, cnt, nc, f, l :=  false, false, 0, 0, rune(0), rune(0)
	var data []Tag
	*src+=" "
	for _,c:=range *src {
		if c=='>' {
			cnt+=1
			sn:= bool(lastTag==currentTag) 
			if f!='/' && sn {nc+=1}
			if f=='/' && sn {nc-=1}
			if cnt==1 && (l =='/' || f=='!' || f=='?')  {
				content, Props, cnt =  "", Props+" ", 0
				data = append(data,Tag{Name:currentTag,Props:getProps(Props),Type:1})
				Props =  ""
			}
			if cnt>2 && nc==0 {
				content, Props = content+">", Props+" "
				data = append(data,Tag{Name:currentTag,Props:getProps(Props),Type:3, Children:Parse(&content)})
				content, Props, cnt =  "", "", 0
			}
			if cnt==2 && nc==0 {
				content, Props = content[:len(content) - len(lastTag) - 1], Props+" "
				data = append(data,Tag{Name:currentTag,InnerText:strip(content),Props:getProps(Props),Type:2})
				content, Props, cnt =  "", "", 0
			}
			if cnt==1 {lastTag, currentTag, nc, n, p  = currentTag, "", 1, false, false; continue}
			currentTag, n,  p  = "", false, false
		}
		if cnt==0 && p {Props+=string(c)}
		if n && c==' ' {p = true}
		if c=='<' {f, n = 0, true}
		if n && !p && c!='/' && c!='<' && c>32 {currentTag+=string(c)}
		if cnt>=1 {content+=string(c)}
		if n && !p && (c=='/' || c=='!' || c=='?') {f = c}
		if c>32 {l = c}
	}
	return data
}

func Stringify(data *[]Tag)string{
	res,lb := "","\n"
	for _,v:= range *data{
		if v.Type==1 {res+="  "+"<"+v.Name+printProps(v.Props)+"/>"+lb}
		if v.Type==2 {res+="  "+"<"+v.Name+printProps(v.Props)+">"+v.InnerText+"</"+v.Name+">"+lb}
		if(v.Type==3){
			res+="  "+"<"+v.Name+printProps(v.Props)+">"+lb
			Children, newChildren := Stringify(&v.Children), "  "  
			for _,c:= range Children {
				newChildren+=string(c)
				if c=='\n' {newChildren+="  "}
			}
			res+=newChildren+"</"+v.Name+">"+lb
		}
	}
	return res
} 