async function upload(event) {
    event.preventDefault();

    const file = document.getElementById("fileSelect").files[0];
    const formData = new FormData();
     
    formData.append("file", file);
    await fetch('http://localhost:3333/form', {method: "POST", body: formData});
}

async function getImports() {
    const imports = await (await fetch('http://localhost:3333/imports', {method: "GET"})).json();
    const rows = imports.map((i, index) => {
        const tr = document.createElement("tr");
        const elementScope = document.createElement("th");
        const elementTransactionsTime = document.createElement("td");
        const elementImportTime = document.createElement("td");
        elementScope.setAttribute("scope", "row")
        elementScope.textContent = index + 1
        elementTransactionsTime.textContent = i.TimeOfTransactions
        elementImportTime.textContent = i.TimeOfImportation
        tr.appendChild(elementScope)
        tr.appendChild(elementTransactionsTime)
        tr.appendChild(elementImportTime)
        return tr
    })

    const table = document.getElementById("imports");
    rows.forEach(row => {
        table.appendChild(row);
    })
}
getImports()