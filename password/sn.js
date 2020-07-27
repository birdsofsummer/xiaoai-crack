// https://bbs.hassbian.com/thread-8860-1-1.html
//https://bbs.hassbian.com/thread-4961-1-1.html


const u=require('utility')
const get_pass=(sn="")=>{
   let b="9C78089F-83C7-3CDC-BCC9-93B378868E7F"
   let c=u.md5(sn+b).slice(0,14)
   console.log(c)
   return c
}

sn="14891/980295402"
get_pass(sn)
// a1914c5e6c45eb
