var wsUri;
var output;
var count;
var ws;

window.addEventListener("load", function(evt) {
    wsUri = "ws://" + window.location.host + "/log/status";

    var print = function(data) {
        message = data.info
        parts = message.split("`")
        d = JSON.parse(parts[1])
        cols = d.cols
        rows = d.rows
        content = ""
        content += "<thead><tr>"
        for (var i = 0; i < cols.length; i++) {
            content += "<th>" + cols[i] + "</th>"
        }
        content += "</tr></thead>"
        content += "<tbody>"
        for (var j = 0; j < rows.length; j++) {
            content += "<tr>"
            for (var k = 0; k < rows[j].length; k++) {
                content += "<td>" + rows[j][k] + "</td>"
            }
            content += "</tr>"
        }
        content += "</tbody>"
        $('#tt').html(content)

        lines = parts[2].split("\n")
        lines.push('Last updated at: ' + new Date(parseInt(data.ts, 10) * 1000).toISOString())
        $('#extra').html("<p>" + lines.join("</p><p>") + "</p>")
    };

    var parseInfo = function(evt) {
        return JSON.parse(evt.data)
    };

    var newSocket = function() {
        ws = new WebSocket(wsUri);
        ws.onopen = function(evt) {
            //print('<span style="color: green;">Connection Open</span>');
        }
        ws.onclose = function(evt) {
            //print('<span style="color: red;">Connection Closed</span>')
            ws = null;
        }
        ws.onmessage = function(evt) {
            print(parseInfo(evt));
        }
        ws.onerror = function(evt) {
            //print('<span style="color: red;">Error: </span>' + parseInfo(evt));
        }
    };

    newSocket()

    $('#btn-go').click(function() {
        if (!ws) {
            newSocket()
        }
        var req = {
            info: $('#tag').val()
        }
        ws.send(JSON.stringify(req))
        return false
    })
})
