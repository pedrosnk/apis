defmodule ApiDash.PageController do
  use ApiDash.Web, :controller

  def index(conn, _params) do
    render conn, "index.html"
  end
end
