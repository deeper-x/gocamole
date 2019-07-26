function getRemoteData(uri){
    fetch(uri, {
        method: 'get'
    })
    .then(response => response.json())
    .then(jsonData => document.getElementById('toPush').innerHTML = jsonData.result)
    .catch(err => {
            console.log(err);
        });
}
    
