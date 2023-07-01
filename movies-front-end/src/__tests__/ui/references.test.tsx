import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import References from "@/pages/admin/references";
import { useRouter } from "next/router";
import { useSession } from "next-auth/react";


jest.mock("next/router", () => ({
  ...(jest.requireActual("next/router") as object),
  useRouter: jest.fn()
}));

jest.mock("next-auth/react", () => ({
  ...(jest.requireActual("next-auth/react") as object),
  useSession: jest.fn()
}));

describe("References component", () => {

  test("renders the component correctly", () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {},
      push: jest.fn()
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    render(<References />);

    // Check if the main heading is rendered
    expect(screen.getByRole("heading", { level: 4 })).toHaveTextContent("The Movie Database Reference");

    // Check if the search button is rendered
    expect(screen.getByRole("button", { name: "Search" })).toBeInTheDocument();
  });

  test("triggers search correctly on button click", async () => {
    render(<References />);

    // Enter a search keyword
    const keywordInput = screen.getByLabelText("Keyword");
    fireEvent.change(keywordInput, { target: { value: "test" } });

    // Click the search button
    const searchButton = screen.getByRole("button", { name: "Search" });
    fireEvent.click(searchButton);

    await waitFor(() => {
      expect(screen.getByText("Test movie 1")).toBeInTheDocument();
    });

  });

  test("triggers search correctly on Enter key press", async () => {
    render(<References />);

    // Enter a search keyword
    const keywordInput = screen.getByLabelText("Keyword");
    fireEvent.change(keywordInput, { target: { value: "test" } });

    // Press Enter key
    fireEvent.keyDown(keywordInput, { key: "Enter" });

    await waitFor(() => {
      expect(screen.getByText("Test movie 1")).toBeInTheDocument();
    });
  });
});