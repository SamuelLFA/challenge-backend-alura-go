async function upload(event) {
    event.preventDefault();

    const file = document.getElementById("fileSelect").files[0];
    const formData = new FormData();
     
    formData.append("file", file);
    await fetch('http://localhost:3333/form', {method: "POST", body: formData});
}