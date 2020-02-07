var apiHost = "localhost:7929";
var apiRouter = "api"
var apiVersion = "v0.0.1";


function getMatch(inp) {
    let query = `http://${apiHost}/${apiRouter}/${apiVersion}/search_v2?req=${inp}`;
    console.log(query);
    (async () => {
        try {
            let response = await fetch(query, {
                headers: {
                    "Accept": "application/json",
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                method: "GET",
            });
            let responseJSON = await response.json();
            if (responseJSON.data.data != null) {
                alert("Получены данные:" + JSON.stringify(responseJSON.data));
            } else {
                alert("Данных не найдено");
            }
        }
        catch (e) {
            console.log("fetch error", e);
        }
    })();
}