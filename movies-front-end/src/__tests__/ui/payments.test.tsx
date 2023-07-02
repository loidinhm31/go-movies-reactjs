import { render, screen, waitFor } from "@testing-library/react";
import PaymentsTable from "@/components/Tables/PaymentsTable";
import userEvent from "@testing-library/user-event";

describe("Payments", () => {
  test("renders table with data", async () => {

    render(<PaymentsTable />);

    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
    });

    // Verify that search input is present
    expect(screen.getByLabelText("Keyword")).toBeInTheDocument();

    // Type a search keyword and trigger search
    const searchInput = screen.getByLabelText("Keyword");
    userEvent.type(searchInput, "movie");

    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
    });
  });

  test("Sort table", async () => {
    render(<PaymentsTable />);

    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
    });

    const typeColumn = screen.getByRole("button", { name: "Type" });
    expect(typeColumn).toBeInTheDocument();

    userEvent.click(typeColumn);

    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
    });
  });
});
