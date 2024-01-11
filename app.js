const filelist = ['images','contact','cloudbag'];

function searchcontent(){
    var searchItem = document.getElementById('searchbar').value.toLowerCase();
    if(filelist.includes(searchItem)){window.location.href = searchItem+".html"}
    else{alert("No matching files found.")}
}

function searchsuggestion(){
    const input = document.getElementById('searchbar');
            const searchTerm = input.value.toLowerCase();
            const searchSuggestions = document.getElementById('suggestions');
            searchSuggestions.innerHTML = ''; // Clear previous suggestions

            if (searchTerm.length === 0) {
                return;
            }

            const matchingFiles = filelist.filter(file => file.toLowerCase().includes(searchTerm));
            if (matchingFiles.length === 0) {
                searchSuggestions.innerHTML = '<li>No matching files found.</li>';
            } else {
                matchingFiles.slice(0, 5).forEach(file => { // Show up to 5 suggestions
                    const listItem = document.createElement('li');
                    const link = document.createElement('a');
                    link.href = file+".html";
                    link.textContent = file;
                    listItem.appendChild(link);
                    searchSuggestions.appendChild(listItem);
                });
            }
        }
