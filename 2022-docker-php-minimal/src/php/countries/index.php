<?php
$cc = $_GET['cc']; // alpha_2 code

function not_found()
{
    // post 0 not allowed
    header("HTTP/1.1 404 Not Found");
    print "<html><body style='text-align:center'><h1>404 Not Found</h1><hr>404 from PHP</body></html>";
    exit();
}

if (!$cc) {
    not_found();
}

mysqli_report(MYSQLI_REPORT_ERROR | MYSQLI_REPORT_STRICT);

$mysqli = mysqli_connect(
    getenv('APP_DB_SERVER'),
    getenv('APP_DB_USER'),
    getenv('APP_DB_PASSWORD'),
    getenv('APP_DB_NAME')
);

mysqli_set_charset($mysqli, 'utf8mb4');

$cc = mysqli_real_escape_string($mysqli, $cc);
$query = "select * from countries where alpha_2 = '$cc'";
$result = mysqli_query($mysqli, $query);

if (mysqli_num_rows($result) < 1) {
    not_found();
}

$content = '';
$row = mysqli_fetch_assoc($result);

$content .= "<table>";
foreach ($row as $key => $val) {
    $content .= "<tr><td>$key</td><td>$val</td></tr>";
}
$content .= "</table>";

$html = <<<EOHTML
<!doctype HTML>
<html lang="en-us" />
<head>
    <title>{$row['en']}</title>
    <meta charset="utf-8" />
</head>
<body>
    <h1>{$row['en']}</h1>
    $content
</body>
</html>
EOHTML;

print $html;
