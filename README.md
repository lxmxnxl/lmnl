![lxmxnxl-banner](https://user-images.githubusercontent.com/96031819/168491238-5141d096-1dcd-41fb-bcd5-29bab702e8bc.jpg)

<small>image credit: https://wallpapercave.com/</small>

# LMNL
<strong>lmnl is a tag-based imageboard in liminal space, written in Go</strong>


## About

lmnl is a tag-based imageboard. lmnl/lxmxnxl is pronounced like 'liminal'. 

This project came from a discussion with a friend about online social anti-networks with no identities, likes, favorites, or other engagement mechanisms. 

Through a conversation with another friend, a challenge was created to write this software in the span of 1 laptop charge. The site was completed with over 50% to spare, the rest is being used toward open sourcing and deploying the site as it will stand on https://www.lxmxnxl.com

## How to deploy
**1. clone it**
```
git clone https://github.com/lxmxnxl/lmnl
```
**2. build it**
```
cd lmnl
go mod init lmnl
go mod tidy 
go build
```
**3. run it**
```
./lmnl
```
**4. optional:**
- nginx reverse proxy to send port :80 traffic to :8666

## Screenshot

![image](https://user-images.githubusercontent.com/96031819/168922056-5d98c05a-929d-4326-8c11-7f71b4241297.png)

## Contributing

- keep it boring, keep it simple

- please don't add dependencies, emoji, or fancy js

## Thanks

- El Eadi Argeebee
- IBMCD
- Monster Energy Ultra Fiesta

## License

the mit license (mit)

copyright © 2022 the lmnl authors

permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “software”), to deal in the software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the software, and to permit persons to whom the software is furnished to do so, subject to the following conditions:

- the above copyright notice and this permission notice shall be included in all copies or substantial portions of the software.

- the software is provided “as is”, without warranty of any kind, express or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose and noninfringement. in no event shall the authors or copyright holders be liable for any claim, damages or other liability, whether in an action of contract, tort or otherwise, arising from, out of or in connection with the software or the use or other dealings in the software.

