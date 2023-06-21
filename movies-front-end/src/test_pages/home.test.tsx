import { render, screen } from "@testing-library/react";
import Home from "src/pages/index";
import userEvent from "@testing-library/user-event";

describe("Home page", () => {
    it("renders the title", () => {
        render(<Home />);
        const titleElement = screen.getByText("Find a movie to watch tonight!");
        expect(titleElement).toBeInTheDocument();
    });

    it("renders the link to movies page", () => {
        render(<Home />);

        const imageElement = screen.getByAltText("movie tickets");

        const linkElement = screen.getByRole("link", { name: "movie tickets" });

        expect(linkElement).toHaveAttribute("href", "/movies");
    });
});
