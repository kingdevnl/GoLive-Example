<div style="overflow-y: scroll; overscroll-behavior-y: contain; scroll-snap-type: y proximity; height: 100%; padding-left: 5px; padding-right: 5px">
    {{range $message := .messages}}
        <div>
            <span class="name">{{$message.User}}.</span>
            <span class="time">{{$message.Date}}</span>
            <p>{{$message.Message}}</p>
        </div>
        <hr>
    {{end}}
</div>