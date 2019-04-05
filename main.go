package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var ipChangeTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style type="text/css">

        .myButton {
            -moz-box-shadow:inset 0px 1px 0px 0px #cf866c;
            -webkit-box-shadow:inset 0px 1px 0px 0px #cf866c;
            box-shadow:inset 0px 1px 0px 0px #cf866c;
            background:-webkit-gradient(linear, left top, left bottom, color-stop(0.05, #d0451b), color-stop(1, #bc3315));
            background:-moz-linear-gradient(top, #d0451b 5%, #bc3315 100%);
            background:-webkit-linear-gradient(top, #d0451b 5%, #bc3315 100%);
            background:-o-linear-gradient(top, #d0451b 5%, #bc3315 100%);
            background:-ms-linear-gradient(top, #d0451b 5%, #bc3315 100%);
            background:linear-gradient(to bottom, #d0451b 5%, #bc3315 100%);
            filter:progid:DXImageTransform.Microsoft.gradient(startColorstr='#d0451b', endColorstr='#bc3315',GradientType=0);
            background-color:#d0451b;
            -moz-border-radius:3px;
            -webkit-border-radius:3px;
            border-radius:3px;
            border:1px solid #942911;
            display:inline-block;
            cursor:pointer;
            color:#ffffff;
            font-family:Arial;
            font-size:15px;
            padding:6px 24px;
            text-decoration:none;
            text-shadow:0px 1px 0px #854629;
        }
        .myButton:hover {
            background:-webkit-gradient(linear, left top, left bottom, color-stop(0.05, #bc3315), color-stop(1, #d0451b));
            background:-moz-linear-gradient(top, #bc3315 5%, #d0451b 100%);
            background:-webkit-linear-gradient(top, #bc3315 5%, #d0451b 100%);
            background:-o-linear-gradient(top, #bc3315 5%, #d0451b 100%);
            background:-ms-linear-gradient(top, #bc3315 5%, #d0451b 100%);
            background:linear-gradient(to bottom, #bc3315 5%, #d0451b 100%);
            filter:progid:DXImageTransform.Microsoft.gradient(startColorstr='#bc3315', endColorstr='#d0451b',GradientType=0);
            background-color:#bc3315;
        }
        .myButton:active {
            position:relative;
            top:1px;
        }


        .form-style-7{
            max-width:400px;
            margin:50px auto;
            background:#fff;
            border-radius:2px;
            padding:20px;
            font-family: Georgia, "Times New Roman", Times, serif;
        }
        .form-style-7-h2{
            display: block;
            text-align: center;
            padding: 0;
            margin: 0px 0px 20px 0px;
            color: #5C5C5C;
            font-size:x-large;
        }
        .form-style-7 ul{
            list-style:none;
            padding:0;
            margin:0;
        }
        .form-style-7 li{
            display: block;
            padding: 9px;
            border:1px solid #DDDDDD;
            margin-bottom: 30px;
            border-radius: 3px;
        }
        .form-style-7 li:last-child{
            border:none;
            margin-bottom: 0px;
            text-align: center;
        }
        .form-style-7 li > label{
            display: block;
            float: left;
            margin-top: -19px;
            background: #FFFFFF;
            height: 14px;
            padding: 2px 5px 2px 5px;
            color: #B9B9B9;
            font-size: 14px;
            overflow: hidden;
            font-family: Arial, Helvetica, sans-serif;
        }
        .form-style-7 input[type="text"],
        .form-style-7 input[type="date"],
        .form-style-7 input[type="datetime"],
        .form-style-7 input[type="email"],
        .form-style-7 input[type="number"],
        .form-style-7 input[type="search"],
        .form-style-7 input[type="time"],
        .form-style-7 input[type="url"],
        .form-style-7 input[type="password"],
        .form-style-7 textarea,
        .form-style-7 select
        {
            box-sizing: border-box;
            -webkit-box-sizing: border-box;
            -moz-box-sizing: border-box;
            width: 100%;
            display: block;
            outline: none;
            border: none;
            height: 40px;
            line-height: 25px;
            font-size: 16px;
            padding: 0;
            font-family: Georgia, "Times New Roman", Times, serif;
        }
        .form-style-7 input[type="text"]:focus,
        .form-style-7 input[type="date"]:focus,
        .form-style-7 input[type="datetime"]:focus,
        .form-style-7 input[type="email"]:focus,
        .form-style-7 input[type="number"]:focus,
        .form-style-7 input[type="search"]:focus,
        .form-style-7 input[type="time"]:focus,
        .form-style-7 input[type="url"]:focus,
        .form-style-7 input[type="password"]:focus,
        .form-style-7 textarea:focus,
        .form-style-7 select:focus
        {
        }
        .form-style-7 li > span{
            background: #F3F3F3;
            display: block;
            padding: 3px;
            margin: 0 -9px -9px -9px;
            text-align: center;
            color: #C0C0C0;
            font-family: Arial, Helvetica, sans-serif;
            font-size: 11px;
        }
        .form-style-7 textarea{
            resize:none;
        }
        .form-style-7 input[type="submit"],
        .form-style-7 input[type="button"]{
            background: #2471FF;
            border: none;
            padding: 10px 20px 10px 20px;
            border-bottom: 3px solid #5994FF;
            border-radius: 3px;
            color: #D2E2FF;
            width: 100%;
        }
        .form-style-7 input[type="submit"]:hover,
        .form-style-7 input[type="button"]:hover{
            background: #6B9FFF;
            color:#fff;
        }
        #lined {
            display: block;
            margin: auto;
            width: 30%;
            height: 25rem;

            font-size: 20px;
            line-height: 20px;
        }
    </style>
</head>
<body>

<h2 class="form-style-7-h2">Armbian ip settings
    <input class="myButton" type="button" value="reboot" onclick="location.href='/reboot';">
    
</h2>
<textarea id="lined" cols="30" rows="10" readonly>{{ range . }}
        {{ . }}
    {{ end }}
</textarea>
<form class="form-style-7" action="/" method=post>
    <ul>
        <li>
            <label for="ipAddress">Ip</label>
            <input type="text" required pattern="((^|\.)((25[0-5])|(2[0-4]\d)|(1\d\d)|([1-9]?\d))){4}$" id="ipAddress" placeholder="Ip adresi giriniz" name="ipAddress">
            <span>Ip adresi giriniz</span>
        </li>
        <li>
            <label for="gateway">Gateway</label>
            <input type="text" required pattern="((^|\.)((25[0-5])|(2[0-4]\d)|(1\d\d)|([1-9]?\d))){4}$" id="gateway" placeholder="gateway adresi giriniz" name="gateway">
            <span>Gateway adresi giriniz</span>
        </li>
        <li>
            <input type="submit" value="Gönder" >
        </li>
    </ul>
</form>

</body>
</html>

`

func handler(w http.ResponseWriter, r *http.Request) {
	ipSettingsTemplate := template.New("main")
	ipSettingsTemplate, _ = ipSettingsTemplate.Parse(ipChangeTemplate)
	switch r.Method {
	case "GET":
		str := readLine("/etc/network/interfaces")
		err := ipSettingsTemplate.Execute(w, str)
		if err != nil {
			fmt.Println("error template parsing")
		}
	case "POST":
		ipAddress := r.FormValue("ipAddress")
		gateway := r.FormValue("gateway")
		ipSplit := strings.Split(ipAddress, ".")
		gatewaySplit := strings.Split(gateway, ".")
		if ipSplit[2] == gatewaySplit[2] {
			armbianStatic(ipAddress, gateway)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}

	}

}
func rebootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("reboot").Output()
	if err != nil {
		fmt.Println(err, "execCommand reboot")
		w.Write([]byte("false"))
	}
	w.Write([]byte("true"))
}

func main() {
	server := http.Server{
		Addr: ":8888",
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/reboot", rebootHandler)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error http listen server")
	}
}

func armbianStatic(staticIp, gateway string) {
	ipSplit := strings.Split(staticIp, ".")
	var subNet string
	if len(ipSplit) == 4 {
		subNet = ipSplit[2]
	}
	file, err := os.Create("/etc/network/interfaces")
	if err != nil {
		fmt.Println("error cannot create file" + err.Error())
	}

	_, errFileWrite := file.WriteString(`
source /etc/network/interfaces.d/*
# Network is managed by Network manager
auto lo eth0
iface lo inet loopback

auto eth0
#allow-hotplug eth0
#iface eth0 inet dhcp
iface eth0 inet static
address ` + staticIp + `
netmask 255.255.255.0
broadcast 192.168.` + subNet + `.255
network 192.168.` + subNet + `.0
address ` + gateway + `
dns-nameservers 8.8.8.8 8.8.4.4`)
	if errFileWrite != nil {
		fmt.Println(errFileWrite)
		file.Close()
		return
	}
	errFileWrite = file.Close()
	if errFileWrite != nil {
		fmt.Println("error " + errFileWrite.Error())
		return
	}
}

func readLine(path string) []string {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var str []string
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	return str
}
