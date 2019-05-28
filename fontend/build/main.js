console.log('hellojs')
const host_home = "home-service"
const host_car = "car-service"
// const host_home = "localhost"
// const host_car = "localhost"
const port_home = 9000;
const port_car = 9001;
fetch(`http://${host_home}:${port_home}/`).then(res=>res.json()).then(data=>{
    // console.log(data)
    console.log(data)
    const sel = document.getElementById('home-service');
   
    for (const prop in data) {
        let opt = document.createElement('li');
        let text = prop + ': '+ data[prop]
        opt.appendChild(document.createTextNode(text))
        sel.appendChild(opt); 
    }
// Location: "Thanland"
// Price: 12236222
// Typehome: "single"
    
})
fetch(`http://${host_car}:${port_car}/`).then(res=>res.json()).then(data=>{
    // console.log(data)
    console.log(data)
    const sel = document.getElementById('car-service');
   
    for (const prop in data) {
        let opt = document.createElement('li');
        let text = prop + ': '+ data[prop]
        opt.appendChild(document.createTextNode(text))
        sel.appendChild(opt); 
    }
// Location: "Thanland"
// Price: 12236222
// Typehome: "single"
    
})
document.getElementsByClassName("title")[0].innerHTML = "Hello My servervice"