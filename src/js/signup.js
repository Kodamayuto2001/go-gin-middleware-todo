function handleEvent(event) {
    let name = document.querySelector('#name').value;
    let email = document.querySelector('#email').value;
    let password = document.querySelector('#password').value;

    let data = JSON.stringify({ 
        name: name, 
        email: email, 
        password: password 
    });
    console.log(data)

    let client = new XMLHttpRequest();
    client.open('POST', 'http://localhost:3000/api/v1/users/add');
    client.setRequestHeader('Content-Type', 'application/json');
    client.send(data);
}

function main() {
    button = document.querySelector('button');

    button.addEventListener('click', handleEvent);
}

main();
