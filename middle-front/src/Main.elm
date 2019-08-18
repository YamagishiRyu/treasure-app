module Main exposing (main)

import Browser
import Browser.Navigation as Nav
import Debug
import Html exposing (..)
import Html.Attributes exposing (..)
import Page.RepositoryList as RepositoryList
import Route exposing (Route(..), urlToRoute)
import Url



-- MAIN


main : Program () Model Msg
main =
    Browser.application
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        , onUrlChange = UrlChanged
        , onUrlRequest = LinkClicked
        }



-- MODEL


type alias Model =
    { key : Nav.Key
    , route : Route
    , repositoryList : RepositoryList.Model
    }


init : () -> Url.Url -> Nav.Key -> ( Model, Cmd Msg )
init flags url key =
    ( { key = key
      , route = urlToRoute url
      , repositoryList = { repositories = [], status = RepositoryList.Loading }
      }
    , Cmd.none
    )


type Msg
    = LinkClicked Browser.UrlRequest
    | UrlChanged Url.Url
    | RepositoryListMsg RepositoryList.Msg


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        LinkClicked urlRequest ->
            case urlRequest of
                Browser.Internal url ->
                    ( model, Nav.pushUrl model.key (Url.toString url) )

                Browser.External href ->
                    ( model, Nav.load href )

        UrlChanged url ->
            ( { model | route = urlToRoute url }
            , Cmd.none
            )

        RepositoryListMsg msg_ ->
            let
                ( m_, cmd ) =
                    RepositoryList.update msg_ model.repositoryList
            in
            ( { model | repositoryList = m_ }, Cmd.map RepositoryListMsg cmd )


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none



-- VIEW


view : Model -> Browser.Document Msg
view model =
    { title = "Markdown Search"
    , body =
        [ text "header"
        , viewLink model
        ]
    }


viewLink : Model -> Html Msg
viewLink model =
    case model.route of
        Index ->
            RepositoryList.view
                model.repositoryList
                |> Html.map RepositoryListMsg

        Show id ->
            div []
                [ text ("Show" ++ String.fromInt id)
                , a [ href "/repositories/index" ] [ text "index" ]
                ]
