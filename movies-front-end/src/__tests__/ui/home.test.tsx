import { render, screen } from "@testing-library/react";
import Home from "@/pages";

describe("Home page", () => {
  test("renders the title", () => {
    render(<Home />);
    const titleElement = screen.getByText("Find a movie to watch tonight!");
    expect(titleElement).toBeInTheDocument();
  });

  test("renders the link to movies page", () => {
    render(<Home />);

    const imageElement = screen.getByAltText("movie tickets");

    const linkElement = screen.getByRole("link", { name: "movie tickets" });

    expect(linkElement).toHaveAttribute("href", "/movies");
  });
});
