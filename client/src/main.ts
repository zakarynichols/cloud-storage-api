function app() {
  const formData = new FormData();

  const inputEl = document.querySelector("input");
  inputEl?.addEventListener("change", (e) => {
    if (!(e.target instanceof HTMLInputElement))
      throw new Error("Event target is not an input element.");

    e.preventDefault();
    appendFormData(e.target, formData);
  });

  const formEl = document.querySelector("form");
  formEl?.addEventListener("submit", async (ev) => {
    ev.preventDefault();
    await saveFile(formData);
  });
}

app();

async function saveFile(formData: FormData) {
  try {
    const response = await fetch("http://localhost:8080/upload", {
      method: "POST",
      body: formData,
    });
    console.log(response);
  } catch (err: unknown) {
    console.error(err);
  }
}

function appendFormData(target: HTMLInputElement, formData: FormData) {
  const file = target.files[0];
  formData.append("upload", file);
}
