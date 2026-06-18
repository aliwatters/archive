<?php
$id = intval($_GET['id'], 10);

if (!$id) {
    // post 0 not allowed
    header("HTTP/1.1 404 Not Found");
    print "<html><body style='text-align:center'><h1>404 Not Found</h1><hr>404 from PHP</body></html>";
    exit();
}

$html = <<<EOHTML
<!doctype HTML>
<html lang="en-us" />
<head>
    <title>POST #{$id}</title>
    <meta charset="utf-8" />
</head>
<body>
    <h1>POST #{$id}</h1>
    Blah, blah blah, foo bar.
</body>
</html>
EOHTML;

print $html;
