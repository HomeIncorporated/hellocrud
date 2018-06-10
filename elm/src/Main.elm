module Main exposing (..)

import Html exposing (Html, a, br, button, code, div, h1, hr, img, li, nav, span, table, tbody, td, text, th, thead, tr, ul)
import Html.Attributes exposing (attribute, class, href, id, style, type_)
import Navigation exposing (Location)
import Platform.Cmd as Cmd exposing (map)
import Types exposing (CrudRoute, Flags, Model, CommentModel, Route(..))
import UrlParser exposing ((</>), map, s)
import Comment


---- INIT ----


init : Flags -> Location -> ( Model, Cmd Msg )
init flags location =
    let
        ( commentModel, _ ) =
            Comment.init

        currentRoute =
            routeFromLocation location

        model =
            Model currentRoute
                commentModel
    in
        updateModelByLocation model currentRoute



---- UPDATE ----


type Msg
    = OnLocationChange Location
    | CommentMessage Comment.Msg


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        OnLocationChange location ->
            let
                newroute =
                    routeFromLocation location
            in
                updateModelByLocation { model | route = newroute } newroute

        CommentMessage m ->
            let
                ( comment, cmd ) =
                    Comment.update m model.comment
            in
                ( { model | comment = comment, route = CommentPage comment.route }, Cmd.map CommentMessage cmd )



---- VIEW ----


view : Model -> Html Msg
view model =
    let
        innerContent =
            case model.route of
                NotFoundPage ->
                    text "Not found"

                HomePage ->
                    div [] []

                CommentPage crudroute ->
                    div []
                        [ Html.map CommentMessage (Comment.view crudroute model.comment)
                        ]
    in
        -- Render a common navbar + route specific `innerContent`
        div []
            [ nav [ class "navbar navbar-expand-md navbar-dark bg-dark fixed-top" ]
                [ a [ class "navbar-brand", href "/" ]
                    [ text "Home" ]
                , button [ class "navbar-toggler", attribute "data-target" "#navbarSupportedContent", attribute "data-toggle" "collapse", type_ "button" ]
                    [ span [ class "navbar-toggler-icon" ]
                        []
                    , text "  "
                    ]
                , div [ class "collapse navbar-collapse", id "navbarSupportedContent" ]
                    [ ul [ class "navbar-nav mr-auto" ]
                        [ text ""
                        , li [ class "nav-item" ]
                            [ a [ class "nav-link", href "/comment" ]
                                [ text "Comments" ]
                            ]
                        ]
                    ]
                ]
            , innerContent
            ]



---- PROGRAM ----


main : Program Flags Model Msg
main =
    Navigation.programWithFlags OnLocationChange
        { view = view
        , init = init
        , update = update
        , subscriptions = always Sub.none
        }



---- ROUTING ----


routes : UrlParser.Parser (Route -> a) a
routes =
    UrlParser.oneOf
        [ UrlParser.map HomePage UrlParser.top
        , UrlParser.map CommentPage (s "comment" </> Comment.routes)
        ]



-- Given the browser URL, what Route are we?


routeFromLocation : Navigation.Location -> Route
routeFromLocation location =
    case UrlParser.parsePath routes location of
        Just route ->
            route

        Nothing ->
            NotFoundPage



-- Given the browser URL, what is our updated Model? what is our next Cmd Msg?
-- e.g. after authentication, model is `{ model | loggedIn = True }` and Cmd is `Http.get "/newsfeed.json"`


updateModelByLocation : Model -> Route -> ( Model, Cmd Msg )
updateModelByLocation model route =
    case route of
        NotFoundPage ->
            ( model, Cmd.none )

        HomePage ->
            ( model, Cmd.none )

        CommentPage crudroute ->
            let
                ( comment, cmd ) =
                    Comment.updateModelByLocation model.comment crudroute
            in
                ( { model | comment = comment }, Cmd.map CommentMessage cmd )
