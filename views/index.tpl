<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Bootstrap CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">


</head>

<style>
    .chat-container {
        border: 1px solid #ccc;
        margin-top: 50px;
        padding: 25px;
        height: 400px;
    }

    .chat-container .name {
        font-weight: bold;
    }

    .chat-container .time {
        float: right;
        opacity: 0.7;
    }
    .input {
        margin-top: 20px;
        text-align: center
    }
</style>

<body>

<div class="chat-container">
    {{live "messages"}}
</div>
{{live "inputs"}}



<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM"
        crossorigin="anonymous"></script>

</body>

<script src="/js/golive.js"></script>
</html>