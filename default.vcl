vcl 4.0;

backend default {
    .host = "127.0.0.1";
    .port = "8081";
}

sub vcl_recv {
    if (req.method != "GET" && req.method != "HEAD") {
        return (pass);
    }

    return (hash);
}

sub vcl_backend_response {
    if (beresp.http.Surrogate-Key) {
        set beresp.http.X-Surrogate-Key = beresp.http.Surrogate-Key;
    }

    if (bereq.method == "POST" || bereq.method == "PUT" || bereq.method == "DELETE") {
        if (beresp.http.Surrogate-Key) {
            ban("obj.http.X-Surrogate-Key ~ " + beresp.http.Surrogate-Key);
        }
        set beresp.ttl = 0s;
        set beresp.uncacheable = true;
        return (deliver);
    }

    if (beresp.status == 200 && (bereq.method == "GET" || bereq.method == "HEAD")) {
        if (bereq.http.Authorization) {
            set beresp.http.Vary = "Authorization";
        }
        
        if (!beresp.http.Cache-Control) {
            set beresp.ttl = 1h;
        }
    } else {
        set beresp.ttl = 0s;
    }
    
    return (deliver);
}

sub vcl_deliver {
    set resp.http.Via = "login-cache";

    if (obj.hits > 0) {
        set resp.http.X-Cache = "HIT";
    } else {
        set resp.http.X-Cache = "MISS";
    }

    return (deliver);
}