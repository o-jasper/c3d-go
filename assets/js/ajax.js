
// general framework for ajax calls. _callback functions are in callback.js
// should probably just use websockets though 

function new_request_obj(){
    if (window.XMLHttpRequest)
        return new XMLHttpRequest();
    else
        return new ActiveXObject("Microsoft.XMLHTTP");
}

function request_callback(xmlhttp, _func, args){
    xmlhttp.onreadystatechange=function(){
        if (xmlhttp.readyState==4 && xmlhttp.status==200){
		args.unshift(xmlhttp);
		_func.apply(this, args);
        }
    }
}

function make_request(xmlhttp, method, path, async, params){
    xmlhttp.open(method, path, async);
    xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    var s = "", i = 0;
    for (var k in params){
	if (i > 0)
  	  s += "&";
	s += k+"="+encodeURIComponent(params[k]);
	i++;
    }
    //xmlhttp.setRequestHeader("Content-length", s.length); // important?
    xmlhttp.send(s);
}

function _send_tx(a, b){

}

function send_tx(){
    var a = document.forms['transact_form'].getElementsByTagName('input');
    var to = a[0].value;
    var value = a[1].value;
    var gas = a[2];
    var gasP = a[3];
    var from = a[4];

    var args = {};
    for (i=0;i<a.length;i++){
        args[a[i].name] = a[i].value;
    }

    xmlhttp = new_request_obj();
    request_callback(xmlhttp, _send_tx, [name, content]);
    make_request(xmlhttp, "POST", "/transact", true, args);
    return false;
}

