## parsing the rss data i got | GUID

the **guid** field discriminates between reviews and lists, so maybe select only the ones that have the review flag.

```go
    if strings.Contains(GUID, "review"){
        //add to an array of elements maybe?
    }
```

## description parsing

the user's review is in the description field, but the thing is, we also have the poster's image in the description field for some reason
maybe i can filter out by fields in **p tags**?

## learn HTMX

please for the love of God actually try and learn some good developer practices for writing not only HTMX but in general when making the backend - frontend.

## RSS return

Title in **filmTitle.Value**

```go
    for _, item := range feed.Items{
        //This is the title, everything under extensions["Letterboxd"] contains film info
        item.Extensions["letterboxd"]["filmTitle"]
    }
```
