const url = 'https://nn88u8jm0d.execute-api.us-west-2.amazonaws.com/Prod/counter/'

fetch(url, { method: 'POST' })
    .then(response => response.json())
    .then(json => {
        document.querySelector(".counter").innerHTML = Number(json.value);
    });
