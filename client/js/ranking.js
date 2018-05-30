// var x = {"x":["abc", "321"], "z": ["pox", "321", "any"]}


function generate(results) {
    var arr = [];

    console.log(results);

    Object.keys(results).forEach(function(username) {
        var score = 0;
        var words = results[username];

        words.forEach(function (word) {
            score += calculateScore(word, username, results)
        });

        arr.push({ name: username, totalWords: words.length, score: score})
    });

    return arr.sort(compareScore)
}

function calculateScore(word, username, results) {
    var score = word.length;
    Object.keys(results).forEach(function(key) {
        if (username === key) {
            return
        }
        var words = results[key];

        if (words && words.indexOf(word) != -1) {
            score = 0;
        }
    });

    return score
}

function compareScore(a, b) {
    return b.score - a.score
}