import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import Search from "@/pages/search";
import userEvent from "@testing-library/user-event";
import { SearchField } from "@/components/Search/SearchMovie/SearchField";
import { SearchString } from "@/components/Search/SearchMovie/SearchString";
import { SearchRange } from "@/components/Search/SearchMovie/SearchRange";
import SearchOptions from "@/components/Search/SearchMovie/SearchOptions";
import { SearchGenre } from "@/components/Search/SearchMovie/SearchGenre";
import { SearchDate } from "@/components/Search/SearchMovie/SearchDate";

describe("Search component", () => {
  it("should fetch and display search results", async () => {
    render(<Search />);

    expect(screen.getByRole("radio", { name: "Both" })).toBeChecked();

    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
      expect(screen.getAllByRole("columnheader")).toHaveLength(7);
      expect(screen.getAllByRole("link")).toHaveLength(7);
    });
  });
});

describe("Search Field", () => {
  test("should render correctly", () => {
    const setIsClickSearch = jest.fn();
    const setSearchRequest = jest.fn();
    const fieldDataMap = new Map();
    const setFieldDataMap = jest.fn();

    render(
      <SearchField
        setIsClickSearch={setIsClickSearch}
        setSearchRequest={setSearchRequest}
        fieldDataMap={fieldDataMap}
        setFieldDataMap={setFieldDataMap}
      />
    );

    // Simulate selecting a type
    userEvent.click(screen.getByRole("radio", { name: "Both" }));

    // Simulate clicking the search button
    userEvent.click(screen.getByRole("button", { name: "Search" }));

    expect(setIsClickSearch).toHaveBeenCalledWith(true);
    expect(setSearchRequest).toHaveBeenCalledWith({
      filters: [
        {
          field: "type_code",
          operator: "and",
          def: {
            type: "string",
            values: [""]
          }
        }
      ]
    });
  });
});

describe("SearchString component", () => {
  test("should render correctly", () => {
    const label = "Title";
    const field = "title";
    const defType = "string";
    const handleStringField = jest.fn();

    render(
      <SearchString
        label={label}
        field={field}
        defType={defType}
        handleStringField={handleStringField}
      />
    );

    // Assert that the components are rendered correctly
    expect(screen.getByText(label)).toBeInTheDocument();
    expect(screen.getByLabelText("Operator")).toBeInTheDocument();
    expect(screen.getByLabelText("Field")).toBeInTheDocument();
    expect(screen.getByLabelText("Value")).toBeInTheDocument();
  });

  test("should handle operator select change", () => {
    const label = "Title";
    const field = "title";
    const defType = "string";
    const handleStringField = jest.fn();

    render(
      <SearchString
        label={label}
        field={field}
        defType={defType}
        handleStringField={handleStringField}
      />
    );

    // Simulate selecting an operator
    userEvent.click(screen.getByLabelText("Operator"));
    userEvent.click(screen.getByText("AND"));

    // Assert that the handleStringField function is called with the correct arguments
    expect(handleStringField).toHaveBeenCalledWith(field, "and", "operator", defType);
  });

  test("should handle value input change", () => {
    const label = "Title";
    const field = "title";
    const defType = "string";
    const handleStringField = jest.fn();

    render(
      <SearchString
        label={label}
        field={field}
        defType={defType}
        handleStringField={handleStringField}
      />
    );

    // Simulate entering a value in the input field
    const valueInput = screen.getByPlaceholderText("Title");
    userEvent.type(valueInput, "example{enter}");

    // Assert that the handleStringField function is called with the correct arguments
    expect(handleStringField).toHaveBeenCalledWith(field, ["example"], "def", defType);
  });
});

describe("SearchRange component", () => {
  test("should render correctly", () => {
    const label = "Price";
    const field = "price";
    const defType = "number";
    const min = 0;
    const max = 500;
    const step = 1;
    const handleRangeField = jest.fn();

    render(
      <SearchRange
        label={label}
        field={field}
        defType={defType}
        min={min}
        max={max}
        step={step}
        handleRangeField={handleRangeField}
      />
    );

    // Assert that the components are rendered correctly
    expect(screen.getByText(label)).toBeInTheDocument();
    expect(screen.getByLabelText("Operator")).toBeInTheDocument();
    expect(screen.getByLabelText("Field")).toBeInTheDocument();
    expect(screen.getByTestId("slider")).toBeInTheDocument();
  });

  test("should handle operator select change", () => {
    const label = "Price";
    const field = "price";
    const defType = "number";
    const min = 0;
    const max = 500;
    const step = 1;
    const handleRangeField = jest.fn();

    render(
      <SearchRange
        label={label}
        field={field}
        defType={defType}
        min={min}
        max={max}
        step={step}
        handleRangeField={handleRangeField}
      />
    );

    // Simulate selecting an operator
    userEvent.click(screen.getByLabelText("Operator"));
    userEvent.click(screen.getByText("AND"));

    // Assert that the handleRangeField function is called with the correct arguments
    expect(handleRangeField).toHaveBeenCalledWith(field, "and", "operator", defType);
  });

  test("should handle slider change", () => {
    const label = "Price";
    const field = "price";
    const defType = "number";
    const min = 0;
    const max = 500;
    const step = 1;
    const handleRangeField = jest.fn();

    const container = render(
      <SearchRange
        label={label}
        field={field}
        defType={defType}
        min={min}
        max={max}
        step={step}
        handleRangeField={handleRangeField}
      />
    );
    // Simulate moving the slider
    const slider = screen.getByTestId("slider");
    const inputTag = container.baseElement.querySelectorAll("input[type=\"range\"]");

    fireEvent.change(inputTag[1], { target: { value: 201 } });

    expect(handleRangeField).toHaveBeenCalledWith(field, [0, 201], "def", defType);
  });
});

describe("SearchOptions", () => {
  test("fetches ratings ands displays them in the Autocomplete component", async () => {
    render(
      <SearchOptions label="MPAA Rating" field="rating" defType="string" handleStringField={jest.fn()} />
    );
    // Open the accordion
    userEvent.click(screen.getByRole("button", { name: "MPAA Rating" }));


    await waitFor(() => {
      userEvent.click(screen.getByPlaceholderText("MPAA Rating"));
    });

    // Check that the ratings are displayed in the Autocomplete component
    expect(screen.getByRole("option", { name: "G" })).toBeInTheDocument();
    expect(screen.getByRole("option", { name: "PG" })).toBeInTheDocument();
    expect(screen.getByRole("option", { name: "PG-13" })).toBeInTheDocument();
    expect(screen.getByRole("option", { name: "R" })).toBeInTheDocument();
    expect(screen.getByRole("option", { name: "NC-17" })).toBeInTheDocument();
  });

  test("calls handleStringField when the operator is selected", async () => {
    const handleStringFieldMock = jest.fn();

    render(
      <SearchOptions label="MPAA Rating" field="rating" defType="string" handleStringField={handleStringFieldMock} />
    );

    // Simulate selecting an operator
    userEvent.click(screen.getByLabelText("Operator"));
    userEvent.click(screen.getByText("AND"));

    // Check that handleStringField was called with the correct arguments
    expect(handleStringFieldMock).toHaveBeenCalledWith("rating", "and", "operator", "string");
  });

  test("calls handleStringField when the option is selected", async () => {
    const handleStringFieldMock = jest.fn();

    render(
      <SearchOptions label="MPAA Rating" field="rating" defType="string" handleStringField={handleStringFieldMock} />
    );

    // Open the accordion
    userEvent.click(screen.getByRole("button", { name: "MPAA Rating" }));

    await waitFor(() => {
      userEvent.click(screen.getByPlaceholderText("MPAA Rating"));
    });

    // Select the rating
    userEvent.click(screen.getByRole("option", { name: "G" }));

    // Check that handleStringField was called with the correct arguments
    expect(handleStringFieldMock).toHaveBeenCalledWith("rating", ["G"], "def", "string");
  });
});

describe("SearchGenre", () => {
  test("fetches genres and displays them in the Autocomplete component", async () => {
    render(<SearchGenre movieType="selectedType" handleStringField={jest.fn()} />);

    // Open the accordion
    userEvent.click(screen.getByRole("button", { name: "Genres" }));

    await waitFor(() => {
      userEvent.click(screen.getByRole("combobox", { name: "Genres" }));
    });

    // Check that the genres are displayed in the Autocomplete component
    expect(screen.getByRole("option", { name: "G1 - MOVIE" })).toBeInTheDocument();
    expect(screen.getByRole("option", { name: "G2 - MOVIE" })).toBeInTheDocument();
    expect(screen.getByRole("option", { name: "G5 - TV" })).toBeInTheDocument();
    expect(screen.getByRole("option", { name: "G7 - TV" })).toBeInTheDocument();
  });

  test("calls handleStringField when the operator is selected", () => {
    const handleStringFieldMock = jest.fn();

    render(<SearchGenre movieType="selectedType" handleStringField={handleStringFieldMock} />);

    // Simulate selecting an operator
    userEvent.click(screen.getByLabelText("Operator"));
    userEvent.click(screen.getByText("AND"));

    // Check that handleStringField was called with the correct arguments
    expect(handleStringFieldMock).toHaveBeenCalledWith("genres", "and", "operator", "string");
  });

  test("calls handleStringField when the genre is selected", async () => {
    const handleStringFieldMock = jest.fn();

    render(<SearchGenre movieType="selectedType" handleStringField={handleStringFieldMock} />);


    // Open the accordion
    userEvent.click(screen.getByRole("button", { name: "Genres" }));

    await waitFor(() => {
      userEvent.click(screen.getByRole("combobox", { name: "Genres" }));
    });

    // Select the rating
    userEvent.click(screen.getByRole("option", { name: "G1 - MOVIE" }));

    // Check that handleStringField was called with the correct arguments
    expect(handleStringFieldMock).toHaveBeenCalledWith("genres", ["G1-MOVIE"], "def", "string");
  });
});

describe("SearchDate", () => {
  test("calls handleDateField when the operator is selected", () => {
    const handleDateFieldMock = jest.fn();

    render(
      <SearchDate label="Release Date" field="releaseDate" defType="string" handleDateField={handleDateFieldMock} />
    );

    // Simulate selecting an operator
    userEvent.click(screen.getByLabelText("Operator"));
    userEvent.click(screen.getByText("AND"));

    // Check that handleDateField was called with the correct arguments
    expect(handleDateFieldMock).toHaveBeenCalledWith("releaseDate", "and", "operator", "string", "");
  });

  test("calls handleDateField when the start date is selected", () => {
    const handleDateFieldMock = jest.fn();

    render(
      <SearchDate label="Release Date" field="releaseDate" defType="string" handleDateField={handleDateFieldMock} />
    );

    // Select the start date
    fireEvent.change(screen.getByLabelText("From"), {
      target: { value: "2023-06-24" },
    });

    // Check that handleDateField was called with the correct arguments
    expect(handleDateFieldMock).toHaveBeenCalledWith("releaseDate", "2023-06-24", "def", "string", "from");
  });

  test("calls handleDateField when the end date is selected", () => {
    const handleDateFieldMock = jest.fn();

    render(
      <SearchDate label="Release Date" field="releaseDate" defType="string" handleDateField={handleDateFieldMock} />
    );

    // Select the end date
    fireEvent.change(screen.getByLabelText("To"), {
      target: { value: "2023-06-30" },
    });

    // Check that handleDateField was called with the correct arguments
    expect(handleDateFieldMock).toHaveBeenCalledWith("releaseDate", "2023-06-30", "def", "string", "to");
  });
});