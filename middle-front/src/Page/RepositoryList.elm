module Page.RepositoryList exposing (..)

import Html exposing (..)
import Html.Events exposing (..)
import Http exposing (..)
import Json.Decode exposing (..)



-- MODEL


type alias Model =
    { repositories : List Repositoy
    , status : Status
    }


type Status
    = Loading
    | Success
    | Failed String


type alias Repositoy =
    { id : Int, name : String, path : String, url : String }


type Msg
    = GetRepositories
    | GotRepositories (Result Http.Error (List Repositoy))


host : String
host =
    "http://localhost:1991/repositories/"


init : () -> ( Model, Cmd Msg )
init _ =
    Debug.log "repository init: "
        ( { repositories = []
          , status = Loading
          }
        , getRepositories
        )



-- UPDATE


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        GetRepositories ->
            ( model
            , getRepositories
            )

        GotRepositories result ->
            case result of
                Ok repositories ->
                    ( { model | repositories = repositories, status = Success }, Cmd.none )

                Err err ->
                    let
                        _ =
                            Debug.log "error:" err
                    in
                    ( { model | status = Failed <| httpErrorToString <| err }, Cmd.none )



-- SUBSCRIPTION


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Model -> Html Msg
view model =
    div []
        [ div [] [ text "Repositoy List" ]
        , button [ onClick GetRepositories ] [ text "fetch" ]
        , div []
            [ viewContent model
            ]
        ]


viewContent model =
    case model.status of
        Success ->
            div []
                (viewRepositoryList model)

        Loading ->
            text "Now Loading..."

        Failed s ->
            text s


viewRepositoryList : Model -> List (Html msg)
viewRepositoryList model =
    List.map
        (\repository ->
            div []
                [ viewRepository repository
                ]
        )
        model.repositories


viewRepository : Repositoy -> Html msg
viewRepository repository =
    div []
        [ text repository.name
        , text <| String.fromInt <| repository.id
        ]



-- API


getRepositories : Cmd Msg
getRepositories =
    Http.request
        { method = "GET"
        , headers =
            [ Http.header "Content-Type" "application/json"
            , Http.header "Accept" "application/json"
            , Http.header "Access-Control-Allow-Origin" "*"
            ]
        , url = host
        , expect = Http.expectJson GotRepositories repositoriesDecoder
        , body = Http.emptyBody
        , timeout = Nothing
        , tracker = Nothing
        }


repositoriesDecoder : Decoder (List Repositoy)
repositoriesDecoder =
    Json.Decode.list repositoryDecoder


repositoryDecoder : Decoder Repositoy
repositoryDecoder =
    Json.Decode.map4 Repositoy
        (Json.Decode.field "id" Json.Decode.int)
        (Json.Decode.field "name" Json.Decode.string)
        (Json.Decode.field "path" Json.Decode.string)
        (Json.Decode.field "url" Json.Decode.string)


httpErrorToString : Http.Error -> String
httpErrorToString err =
    case err of
        BadUrl _ ->
            "BadUrl"

        Timeout ->
            "Timeout"

        NetworkError ->
            "NetworkError"

        BadStatus _ ->
            "BadStatus"

        BadBody s ->
            "BadBody: " ++ s
