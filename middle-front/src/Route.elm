module Route exposing (Route(..), routeParser, urlToRoute)

import Url
import Url.Parser exposing ((</>), Parser, int, map, oneOf, s, string)


type Route
    = Index
    | Show Int


routeParser : Parser (Route -> a) a
routeParser =
    oneOf
        [ map Index (s "repositories")
        , map Show (s "repositories" </> int)
        ]


urlToRoute : Url.Url -> Route
urlToRoute url =
    url
        |> Url.Parser.parse routeParser
        |> Maybe.withDefault Index
