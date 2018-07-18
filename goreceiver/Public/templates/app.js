const DeleteRequest = "http://172.24.0.194:3001/delete/v1?id="
const devDeleteRequest = "http://localhost:3001/delete/v1?id="
function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}  

function deleteFile(url){
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.onerror = function (e) {
        console.error(xhr.statusText);
    };
    
  xhr.onreadystatechange = processRequest;

  async function processRequest(e) {
    if (xhr.readyState == 4 && xhr.status == 200) {
        // time to partay!!!
        await sleep(2000);
        location.reload();
    }
  }
    xhr.send(null);
    //location.reload();
}

var g = document.getElementById('file_list');
for (var i = 0, len = g.children.length; i < len; i++)
{

    (function(index){
        g.children[i].children[1].onclick = function(){
              //console.log(DeleteRequest+index);
              deleteFile(DeleteRequest+index);
        }    
    })(i);

}