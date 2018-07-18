var g = document.getElementById('file_list');
for (var i = 0, len = g.children.length; i < len; i++)
{

    (function(index){
        g.children[i].onclick = function(){
              alert(index)  ;
        }    
    })(i);

}